# MGRun

**MGRun** é um wrapper leve e multiplataforma em Go para execução de comandos do sistema, projetado para simplificar o uso de `exec` com **captura de saída em tempo real**, **callbacks** e **controle seguro de concorrência**.

Ideal para aplicações CLI, ferramentas de automação, instaladores, utilitários desktop e sistemas que precisam acompanhar a execução de comandos enquanto eles ainda estão rodando.

---

## ✨ Recursos

* 🌍 **Multiplataforma**
  Abstrai automaticamente a execução entre:

  * PowerShell (Windows)
  * SH (Linux e macOS)

* 📡 **Streaming em tempo real**
  Receba cada linha de `stdout` e `stderr` via callbacks enquanto o processo executa.

* 🧵 **Thread-safe**
  Leitura concorrente segura das streams e controle confiável do processo.

* 🏠 **Ambiente herdado**

  * Executa comandos a partir do diretório *Home* do usuário
  * Herda variáveis de ambiente do sistema
  * Permite adicionar variáveis customizadas

* 🔁 **Código de saída acessível**
  Obtenha o *exit code* após a finalização do processo.

---

## 📦 Instalação

```bash
go get github.com/mugomes/mgrun
```

---

## 🚀 Exemplo de uso

```go
package main

import (
	"fmt"
	"os"

	"github.com/mugomes/mgrun"
)

func main() {
	sRun := mgrun.New("ls -a")

	// Definir diretório de execução (opcional)
	pathHome, _ := os.UserHomeDir()
	sRun.SetDir(pathHome)

	// Variáveis de ambiente extras (opcional)
	sRun.AddEnv("EXEMPLO", "Valor")

	// Callback para stderr
	sRun.OnStderr(func(line string) {
		fmt.Printf("[STDERR]: %s\n", line)
	})

	// Callback para stdout
	sRun.OnStdout(func(line string) {
		fmt.Printf("[STDOUT]: %s\n", line)
	})

	// Executa o comando
	if err := sRun.Run(); err != nil {
		fmt.Printf("Erro na execução: %v\n", err)
	}

	fmt.Printf("Exit Code: %d\n", sRun.ExitCode())
}
```

---

## ⚙️ Requisitos

* **Go** 1.25.5 ou superior
* **PowerShell** (apenas no Windows)

---

## 🖥️ Sistemas Operacionais

* ✔️ Linux
* ✔️ Windows
* ✔️ macOS (Darwin)

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://mugomes.github.io](https://mugomes.github.io)

📺 [https://youtube.com/@mugomesoficial](https://youtube.com/@mugomesoficial)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgrun/blob/main/LICENSE) license.

All contributions to the MGRun are subject to this license.