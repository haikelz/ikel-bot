package handlers

import (
	"katou-megumi/pkg/utils"

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

	utils.MessageWithReply(s, m, content, logger)
}
