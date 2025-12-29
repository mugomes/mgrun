# MGRun

MGRun é um wrapper leve em Go para exec que simplifica a execução de comandos do shell com suporte nativo a callbacks em tempo real para as saídas de sistema.

🚀 **Funcionalidades**

Multiplataforma: Abstrai a execução entre PowerShell (Windows) e SH (Linux/macOS).

Streaming em Tempo Real: Capture linhas de stdout e stderr via callbacks enquanto o processo ainda está rodando.

Thread-Safe: Gerenciamento seguro de concorrência para leitura de streams e captura de código de saída.

Ambiente Herdado: Executa comandos automaticamente a partir do diretório Home do usuário e herda variáveis de ambiente do sistema.

📦 **Instalação**

```bash
go get github.com/mugomes/mgrun
```

🛠️ **Exemplo de Uso**

```golang
package main

import (
    "fmt"
    "github.com/mugomes/mgrun"
)

func main() {
    go func() {
        sRun := mgrun.New("ls -a")

        // Definir um diretório (Opcional)
        pathHome,_ := os.UserHomeDir()
		sRun.SetDir(pathHome)

        // Variáveis extras (Opcional)
        sRun.AddEnv("EXEMPLO", "Valor")

        // Callback para processar cada linha da saída padrão
        sRun.OnStderr(func(line string) {
            fmt.Printf("[LOG]: %s\n", line)
        })

        sRun.OnStdout(func(line string) {
            fmt.Printf("[LOG]: %s\n", line)
        })

        // Executa e aguarda a conclusão
        if err := sRun.Run(); err != nil {
            fmt.Printf("Erro ou comando falhou: %v\n", err)
        }

        fmt.Printf("Exit Code: %d\n", sRun.ExitCode())
    }()
}
```

## Requerimento

- Go 1.25.5 ou superior
- PowerShell (Windows)

### Sistema Operacional

- Linux
- Windows
- Darwin (macOS)

## License

Copyright (c) 2025 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgrun/blob/main/LICENSE) license.

All contributions to the MGRun are subject to this license.