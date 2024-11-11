package bngipcgo

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func acceptorGoroutine(resolve *BngIpcServer, bgw *sync.WaitGroup) {
	bgw.Add(1)
	go func() {
		// Es wird geprüft ob eine neue Verbindung verfügbar ist
	}()
}

func SetupNewIpcServer(ipcSockName string, onNewProcess OnNewProcessFunction, onError OnErrorFunction, onClosed OnClosedFunction) (*BngIpcServer, error) {
	socketPath := filepath.Join("/tmp", strings.ToLower(ipcSockName)) // Systemspezifischer Pfad

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
		onNewProcess:     onNewProcess,
		onError:          onError,
		onClosed:         onClosed,
		wg:               new(sync.WaitGroup),
	}

	// Das Akzeptieren neuer Prozesse wird gestartet
	bgw := new(sync.WaitGroup)
	acceptorGoroutine(resolve, bgw)

	// Es wird gewartet das der Vorgang erfolgreich war
	bgw.Wait()

	// Die Daten werden zurückgegeben
	return resolve, nil
}
