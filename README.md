# 🗂️ Sortify - Automatic File Organizer

A smart file organization tool that automatically sorts your files into appropriate directories based on their types.

## 🌟 Features

- 📹 Automatically sorts video files (`mp4`, `avi`, `mov`, `mkv`, `flv`, `webm`)
- 🎵 Organizes music files (`mp3`, `wav`, `aac`, `flac`, `m4a`)
- 🖼️ Manages image files (`jpg`, `jpeg`, `png`, `gif`, `webp`, `bmp`, `svg`, `ico`)
- 📄 Sorts documents (`pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`)
- ⚡ Real-time file monitoring and organization
- 🔄 Automatic file renaming for duplicates

## 🚀 Getting Started

### Prerequisites

- Go 1.24.1 or higher
- [fsnotify](github.com/fsnotify/fsnotify) package

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Amul-Thantharate/sortify.git
```

2. Navigate to the project directory:
```bash
cd sortify
```

3. Install dependencies:
```bash
go mod tidy
```

## 💻 Usage

1. Set up your directory paths in the code:
```go
sourceDir = "path/to/watch/directory"
destDirMusic = "path/to/music/directory"
destDirVideo = "path/to/video/directory"
destDirImage = "path/to/image/directory"
destDirDocuments = "path/to/documents/directory"
```

2. Run the application:
```bash
go run main.go
```

## 🛠️ How It Works

- The application uses the `fsnotify` package to monitor a specified directory for new files
- When a new file is created, it checks the file extension
- Based on the extension, the file is automatically moved to the appropriate destination folder
- If a file with the same name exists in the destination, it's automatically renamed with a number suffix

## 📝 Supported File Types

### Images
- jpg, jpeg, png, gif, webp, bmp, svg, ico

### Videos
- mp4, avi, mov, mkv, flv, webm

### Audio
- mp3, wav, aac, flac, m4a

### Documents
- pdf, doc, docx, xls, xlsx, ppt, pptx

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

## ✨ Author

Created by [Amul-Thantharate](https://github.com/Amul-Thantharate)