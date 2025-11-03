# CodingMad — Local AI Coding Assistant

**CodingMad** is a self-contained, local AI development assistant designed to help myself take and refer to notes for programming, and will be able to encompass multiple languages and not generate code. The intent is to help me learn and improve my coding journey.  
Built with **Go**, **SQLite**, **Ollama**, and a lightweight **web interface**, it aims to bring ChatGPT-like coding assistance fully offline and under user control, using local resources and offline capability.

---

## Project Overview

CodingMad is being developed as a **Boot.dev learning capstone** project, showcasing skills in:
- **Go programming**
- **SQL and data persistence**
- **Docker containerization**
- **Local AI model integration (Ollama)**
- **Optional message queue integration (RabbitMQ)**
- **Frontend interface (HTML/CSS/JavaScript)**

The final product will run as a **local web app** — powered by Go on the backend and a browser-based UI for chat, notes, and project management.

---

## Tech Stack

| Layer | Technology | Purpose |
|-------|-------------|----------|
| Backend | **Go (Golang)** | Core logic, API handling, and integration with AI and database |
| AI Engine | **Ollama** | Local LLM inference and prompt processing |
| Database | **SQLite** | Persistent storage for user notes, messages, and session data |
| Frontend | **HTML / CSS / JavaScript** | User interface accessible via localhost |
| Optional Services | **Docker**, **RabbitMQ** | Containerization and async messaging |
| Version Control | **Git & GitHub** | Code management and collaboration |

---

## Development Phases

### **Phase 1 — Project Setup & Repository Initialization** - Started 10/28/2025, Finished 10/28/2025
- Initialize GitHub repository and local Git environment  
- Create `main` and working branches  
- Define project structure and documentation  
- Configure `.gitignore`, `README.md`, and environment setup  

### **Phase 2 — Core Go Backend** - Started 10/28/2025
- Build Go API skeleton  
- Implement configuration loading and logging  
- Connect to SQLite  
- Define base data models and handlers  

### **Phase 3 — Ollama AI Integration** - Pending
- Connect backend to local Ollama API  
- Implement basic prompt/response system  
- Add caching for model responses  

### **Phase 4 — Web Frontend** - Pending
- Create simple HTML/JS interface served by Go  
- Add input/output components for chat and notes  
- Implement API calls between frontend and backend  

### **Phase 5 — Dockerization** - Pending
- Write `Dockerfile` and `docker-compose.yml`  
- Containerize backend, frontend, and database  
- Ensure cross-platform compatibility  

### **Phase 6 — RabbitMQ Integration** - Pending
- Integrate asynchronous message queue for task handling  
- Set up publisher/subscriber model  

### **Phase 7 - Enhancements** - Pending
- Search or tagging for notes
- Export project data to Markdown or JSON
- Integrate authentication or local token
- Add “memory recall” feature (AI references previous notes automatically)
- Build a simple settings page for configuring Ollama model and parameters

### **Phase 8 — Final Polish** - Pending
- Add error handling and logging improvements  
- Improve UI/UX for chat panel  
- Write comprehensive documentation  

---

## Getting Started

```bash
# Clone the repository
git clone https://github.com/YourUsername/CodingMad.git
cd CodingMad

# Create and switch to a working branch
git checkout -b setup/project-structure
