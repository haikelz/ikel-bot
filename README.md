<p align="center">
<img src="pkg/assets/katou-megumi.gif" alt="katou megumi gif" />
</p>

<h1 align="center">Katou Megumi Discord Bot</h1>

<p align="center">A simple Discord Bot</p>

## Features

### Commands

- **`!info`** - Display bot information and command list
- **`!ping`** - Check bot response time
- **`!salam`** - Send greetings to users
- **`!ask <question>`** - Ask questions using Google Gemini AI (2.5 pro)
- **`!jadwalsholat <city>`** - Get prayer times for a specific city
- **`!doa`** - Display daily prayers and supplications
- **`!asmaulhusna`** - Show the 99 Names of Allah (Asmaul Husna)
- **`!jokes`** - Share random jokes
- **`!quote`** - Display inspirational quotes
- **`!editbackground`** - Edit profile background images
- **`!shutdown`** - Turn off the bot

## Getting Started

### Prerequisites

- Go 1.23.1 or higher
- Discord Bot Token
- Google Gemini API Key

### Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd katou-megumi
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:

   ```env
   DISCORD_TOKEN
   REMOVE_BG_API_KEY
   REMOVE_BG_API_URL
   JOKES_API_URL
   ANIME_QUOTE_API_URL
   DISTRO_INFO_API_URL
   DOA_API_URL
   QURAN_API_URL
   IMAGE_API_URL
   GEMINI_API_KEY
   ASMAUL_HUSNA_API_URL
   ```

4. **Run the bot**

   ```bash
   air -c .air.toml
   ```

## Project Structure

```
katou-megumi/
├── cmd/
│   └── app/
│       └── main.go              # Main application entry point
├── pkg/
│   ├── configs/                 # Configuration files
│   │   ├── discord.go           # Discord client configuration
│   │   ├── gemini.go            # Google Gemini AI configuration
│   │   └── zap.go               # Logging configuration
│   ├── entities/                # Data structures
│   │   ├── asmaul_husna.go      # Asmaul Husna data
│   │   ├── doa.go               # Prayer data
│   │   ├── jadwal_sholat.go     # Prayer times data
│   │   ├── joke.go              # Joke data
│   │   └── quote.go             # Quote data
│   ├── handlers/                # Command handlers
│   │   ├── asmaul_husna_handler.go
│   │   ├── background_photo_handler.go
│   │   ├── doa_handler.go
│   │   ├── gemini_handler.go
│   │   ├── info_handler.go
│   │   ├── jadwal_sholat_handler.go
│   │   ├── joke_handler.go
│   │   ├── ping_handler.go
│   │   ├── quote_handler.go
│   │   └── salam_handler.go
│   └── utils/                   # Utility functions
│       ├── date.go              # Date utilities
│       ├── discord.go           # Discord utilities
│       ├── env.go               # Environment loading
│       └── string.go            # String utilities
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies checksum
└── README.md                    # Project documentation
```

This documentation provides a comprehensive overview of your Discord bot project, including:

1. **Feature overview** - All the commands and capabilities
2. **Installation instructions** - Step-by-step setup guide
3. **Project structure** - Clear organization of the codebase
4. **Configuration details** - How to set up the required APIs
5. **Development guidelines** - How to build and contribute
6. **Dependencies** - External services and libraries used

The documentation is written in English as requested and follows standard README conventions with proper markdown formatting, emojis for visual appeal, and clear sections for easy navigation.

## Configuration

### Discord Bot Setup

1. Create a Discord application at [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a bot for your application
3. Copy the bot token and add it to your `.env` file
4. Invite the bot to your server with appropriate permissions

### Google Gemini AI Setup

1. Get an API key from [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Add the API key to your `.env` file

## 🛠️ Development

### Running the Project Using Air

```
air -c .air.toml
```

### Building the Project

```bash
go build -o bin/bot cmd/app/main.go
```

### Running Tests

```bash
go test ./...
```

### Code Structure

- **Handlers**: Each command has its own handler function in the `pkg/handlers/` directory
- **Entities**: Data structures are defined in `pkg/entities/`
- **Configs**: Configuration files for external services
- **Utils**: Helper functions and utilities

## API Dependencies

- **Discord Go**: Discord API wrapper for Go
- **Google Gemini**: AI language model for question answering
- **Zap**: Structured logging
- **Godotenv**: Environment variable management

## License

This project is licensed under the MIT License - see the LICENSE file for details.

**Created by:** [haikelz](https://github.com/haikelz/)

## 🆘 Support

If you encounter any issues or have questions, please open an issue on the GitHub repository.
