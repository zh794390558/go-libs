package main

func SendPcm(data []byte) {
	log.Info("tts udp in bytes: ", len(data))

	dataBuf := bytes.NewBuffer(data)
	dLen := len(data)

	total := 0
	for dataBuf.Len() > 0 {
		buf := dataBuf.Next(1024)
		b, err := udpConn.Write(buf)
		if err != nil {
			log.Errorf("tts udp send err: ", err)

		}
		total += b
	}
	log.Info("tts udp send: ", total, " ", dLen)
}
