package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	text := "Halo, welcome to  Golang Text to Speech.this is template text to speech on golang.thank you"
	err := textToSpeech(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Text has been converted to speech successfully.")
}

func textToSpeech(text string) error {
	// Masukkan API key TTS Anda di sini
	apiKey := "36c0e57f48734fb1b93d37dbc39c24f8"

	// URL layanan TTS
	ttsURL := "https://api.voicerss.org/"

	// Buat permintaan POST untuk mengonversi teks menjadi ucapan
	resp, err := http.PostForm(ttsURL, url.Values{
		"key":    {apiKey},
		"src":    {text},
		"hl":     {"en-us"},          // Bahasa (sesuaikan dengan kebutuhan Anda)
		"r":      {"0"},              // Kecepatan pembacaan (0-10)
		"c":      {"wav"},            // Format output (mp3, wav, ogg, atau acc)
		"f":      {"8khz_8bit_mono"}, // Kualitas suara (opsi lain tersedia)
		"ssml":   {"false"},          // Gunakan SSML (dapat diabaikan)
		"b64":    {"false"},          // Hasil dalam base64 (dapat diabaikan)
		"v":      {"Evie"},           // Suara yang digunakan (opsi lain tersedia)
		"method": {"POST"},           // Metode permintaan (dapat diabaikan)
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Baca dan simpan output audio
	audioData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Simpan output audio ke file
	err = ioutil.WriteFile("output.mp3", audioData, 0644)
	if err != nil {
		return err
	}

	return nil
}
