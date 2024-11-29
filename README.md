```markdown
# DNS Lookup CLI 🔍

A simple and efficient CLI tool built in **Golang** to perform IP lookups for given domain names. The tool fetches the IPv4 addresses associated with one or multiple domain names in a single command.

---

## 🚀 Features

- Fast and reliable domain IP lookup.
- Supports 1–10 domain names in a single execution.
- Lightweight and easy to use.
- Cross-platform support.

---

## 📦 Installation

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/dns-lookup-cli.git
cd dns-lookup-cli
```

### 2. Build and Install
Make sure you have Go installed on your system.
```
bash~go build
go install
```

### 3. Run the CLI
After installation, run the CLI using:
```bash
dnslookup [domain-name]
```

For multiple domains:
```bash
dnslookup example.com google.com github.com
```

---

## 🛠️ Usage

### Single Domain
```bash
dnslookup example.com
```

**Output**:
```
Domain: example.com
IP Addresses:
 - 93.184.216.34
```

### Multiple Domains
```bash
dnslookup example.com google.com
```

**Output**:
```
Domain: example.com
IP Addresses:
 - 93.184.216.34

Domain: google.com
IP Addresses:
 - 142.250.72.46
 - 142.250.72.78
```

---

## 👨‍💻 Requirements

- **Go version**: 1.20 or higher.

---

## 🤝 Contributing

Feel free to open issues or submit pull requests for enhancements or bug fixes. Contributions are welcome! 🎉

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🌟 Show Your Support

If you like this project, please give it a ⭐️ and share it with your peers! 😊
```
https://github.com/Anirudh12345678/dnslookup-cli

