package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func InfoHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := `
# 🤖 Katou Megumi Bot
*My Istri Discord Bot*

---

## 📋 **Daftar Perintah**

### 🔰 **Umum**
- !salam - Memberikan salam kepada pengguna
- !info - Menampilkan informasi tentang bot
- !ask - Mengajukan pertanyaan kepada bot

### 🎨 **Kustomisasi**
- !editbackground - Mengedit background profil
- !sticker - Mengirim sticker

### 🕌 **Islami**
- !jadwalsholat - Menampilkan jadwal sholat
- !doa - Menampilkan doa-doa harian
- !asmaulhusna - Menampilkan Asmaul Husna

### 🎭 **Hiburan**
- !jokes - Menampilkan lelucon
- !animequote - Quote anime random

### 🛠️ **Utilitas**
- !ocr - Membaca teks dari gambar (Optical Character Recognition)

---

## 👨‍💻 **Developer**
**Created by:** [haikelz](https://github.com/haikelz/)

---

*Selamat menggunakan Katou Megumi Bot! 💜*	
`

	_, err := s.ChannelMessageSendReply(m.ChannelID, content, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error("Error sending message", zap.Error(err))
		return
	}
}
