package helper

import "slices"

var categories = map[string][]string{
	"Images": {
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".svg", ".webp",
	},
	"Documents": {
		".pdf", ".doc", ".docx", ".txt", ".odt", ".rtf", ".md",
		".xls", ".xlsx", ".ppt", ".pptx",
	},
	"Fonts": {
		".ttf", "otf", ".woff", ".woff2", ".eot",
	},
	"Videos": {
		".mp4", ".mkv", ".flv", ".avi", ".mov", ".wmv", ".webm",
	},
	"Audio": {
		".mp3", ".wav", ".flac", ".aac", ".ogg", ".m4a",
	},
	"Archives": {
		".zip", ".rar", ".tar", ".tar.gz", ".7z",
	},
	"Executables": {
		".exe", ".msi", ".bin", ".sh", ".bat", ".apk", ".appimage", ".AppImage", ".run", ".elf", ".out",
	},
	"Installers": {
		".deb", ".rpm", ".pkg.tar.zst", ".tgz", ".txz", ".zst", ".gz",
	},
	"Datasets": {
		".csv", ".json", ".xml", ".yaml", ".yml", ".parquet", ".xls", ".xlsx", ".sql",
	},
	"Code": {
		".go", ".py", ".js", ".ts", ".java", ".c", ".cpp", ".cs",
		".php", ".rb", ".swift", ".rs", ".kt", ".html", ".css",
	},
}

func FindCategory(ext string) string {
	for folder, exts := range categories {
		if slices.Contains(exts, ext) {
			return folder
		}
	}
	return "Others"
}
