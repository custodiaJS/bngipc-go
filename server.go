package bngipcgo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func SetupNewIpcServer(ipcSockName string) (*BngIpcServer, error) {
	// Die System Uptime wird ermittelt
	uptime, err := getUptime()
	if err != nil {
		return nil, err
	}

	// Erzeugen des Hashes aus Uptime und SocketName
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%d", uptime) + ipcSockName))
	hashedPath := hex.EncodeToString(hash.Sum(nil))
	socketPath := filepath.Join("/tmp", hashedPath) // Systemspezifischer Pfad
	fmt.Println(socketPath)

	// Prüfen, ob eine Datei unter dem Systemspezifischen Path existiert
	if _, err := os.Stat(socketPath); err == nil {
		// Datei existiert, versuche zu löschen
		if err := os.Remove(socketPath); err != nil {
			return nil, fmt.Errorf("konnte alte Socket-Datei nicht löschen: %w", err)
		}
	}

	// Erstellen eines neuen Unix Sockets
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("konnte Unix Socket nicht erstellen: %w", err)
	}

	// Die Rückgabe wird erzeugt
	resolve := &BngIpcServer{
		listener:         listener,
		processInstances: make([]*BngIpcProcess, 0),
	}

	// Die Daten werden zurückgegeben
	return resolve, nil
}
