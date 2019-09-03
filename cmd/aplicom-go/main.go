package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/einride/aplicom-go/internal/protocol"
	aplicompb "github.com/einride/proto/gen/go/aplicom/v1"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

const (
	dataValidyBitAGPS              = 0x04
	dataValidyBitGPSJamming        = 0x08
	dataValidyBitGPSFixWhenGlonass = 0x10
	dataValidyBitGPSGlonassFix     = 0x20
	dataValidyBitPostFix           = 0x40
	dataValidyBitCurrFix           = 0x80
)

const (
	topic          = "iot-gnss"
	imeiKey        = "imei"
	vehicleTypeKey = "vehicleType"
	dividerMillion = 1000000
)

func main() {
	var portArgument string
	cmd := &cobra.Command{
		Use:   "aplicom-go",
		Short: "Listen in tcp port for Aplicom data and sends parsed data to iot-gnss topic",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Panic("Failed to initialize logging", err)
			}
			port := ":" + portArgument
			lis, err := net.Listen("tcp", port)
			if err != nil {
				logger.Panic("Failed to listen on tcp port", zap.Error(err))
			}
			defer lis.Close()
			project := os.Getenv("GOOGLE_CLOUD_PROJECT")
			if project == "" {
				logger.Panic("GOOGLE_CLOUD_PROJECT environment variable must be set")
			}
			logger.Info("Listen to:", zap.String("port", port))
			for {
				conn, err := lis.Accept()
				if err != nil {
					logger.Warn("Failed to accept listener", zap.Error(err))
				} else {
					go handleConnection(conn, project, logger)
				}
			}
		},
	}
	cmd.Flags().StringVarP(&portArgument,
		"port", "", "5144",
		"Tcp port to listen on.")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleConnection(conn net.Conn, project string, logger *zap.Logger) {
	var packetV3 protocol.PacketV3
	sc := bufio.NewScanner(conn)
	sc.Split(packetV3.ScanAplicomPackets)
	for sc.Scan() {
		if err := packetV3.UnmarshalBinary(sc.Bytes()); err != nil {
			logger.Warn("could not unmarshal binary", zap.Error(err))
		} else if err = sendToPubSub(packetV3, project); err != nil {
			logger.Warn("senToPubSub failed", zap.Error(err))
		}
	}
	if sc.Err() != nil {
		logger.Warn("read from connection failed: ", zap.Error(sc.Err()))
		return
	}
	if err := conn.Close(); err != nil {
		logger.Warn("failed to close connection:", zap.Error(err))
		return
	}
}

func sendToPubSub(packet protocol.PacketV3, project string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		return xerrors.Errorf("failed to create client: %w", err)
	}
	defer client.Close()
	if err := publish(ctx, client, topic, &packet); err != nil {
		return xerrors.Errorf("failed to publish: %w", err)
	}
	return nil
}

func publish(ctx context.Context, client *pubsub.Client, topic string, packet *protocol.PacketV3) error {
	t := client.Topic(topic)
	position := &aplicompb.GNSS{
		GpsTime:          &timestamp.Timestamp{Seconds: int64(packet.GPSTime), Nanos: 0},
		LatitudeDegrees:  float64(packet.Latitude) / dividerMillion,  // translate from millionths of degrees to degrees
		LongitudeDegrees: float64(packet.Longitude) / dividerMillion, // translate from millionths of degrees to degrees
		AltitudeMeters:   float64(packet.GPSAltitude),
		HeadingRadians:   (float64(packet.Heading) * 2 * math.Pi / 180), // Multiply value by 2
		// to get heading in degrees
		NrUsedSatellites:        uint32(packet.NrOfSatellites),
		SpeedKmPerHour:          uint32(packet.Speed),
		IsAgpsValid:             bool((packet.DataValidity & dataValidyBitAGPS) > 0),
		IsGpsJammingDetected:    bool((packet.DataValidity & dataValidyBitGPSJamming) > 0),
		HasGpsFixGlonass:        bool((packet.DataValidity & dataValidyBitGPSFixWhenGlonass) > 0),
		HasGlonassFix:           bool((packet.DataValidity & dataValidyBitGPSGlonassFix) > 0),
		IsGpsPostFixValid:       bool((packet.DataValidity & dataValidyBitPostFix) > 0),
		IsGpsCurrFixValid:       bool((packet.DataValidity & dataValidyBitCurrFix) > 0),
		IsNetworkRegStatusValid: bool((packet.CellInfo.Network & 0x1) > 0),
		Rat:                     uint32((packet.CellInfo.Network >> 1)),
		Lai:                     protocol.ConvertLaiUint32(packet.CellInfo.LAI),
		Lac:                     uint32(packet.CellInfo.LAC),
		Cid:                     packet.CellInfo.CID,
		Rssi:                    int32(packet.CellInfo.RSSI),
		EventId:                 aplicompb.EventId(packet.EventID),
		EventInfo:               uint32(packet.EventInformation),
	}
	positionData, err := proto.Marshal(position)
	if err != nil {
		return err
	}
	result := t.Publish(ctx, &pubsub.Message{
		Attributes: map[string]string{
			imeiKey:        protocol.ConvertToIMEI(packet.UnitIDHighBytes, packet.UnitIDLowBytes),
			vehicleTypeKey: aplicompb.VehicleType_VEHICLE_TYPE_APLICOM.String(),
		},
		Data: positionData,
	})
	_, err = result.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
