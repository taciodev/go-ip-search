# GoIPSearch - Aplicação de Busca de IPs e Nomes de Servidores

Este é um utilitário de linha de comando desenvolvido em Go que permite aos usuários buscar IPs e nomes de servidores da Internet. Ele oferece uma maneira rápida e eficiente de realizar consultas de DNS e obter informações sobre servidores.

## Funcionalidades

- Buscar IPs associados a um nome de domínio
- Buscar nomes de domínio associados a um IP

## Como Usar

1. **Clonar o Repositório:**

   ```bash
   git clone https://github.com/taciodev/go-ip-search.git
   ```

2. **Compilar o Programa:**

   ```bash
   cd go-ip-search
   go build -o goipsearch main.go
   ```
   
3. **Executar o Programa:**

   ```bash
   ./goipsearch
   ```

4. **Interagir com a Aplicação:** Utilize os comandos abaixo para buscar IPs e nomes de servidores:

## Exemplos de Uso

- Para buscar IPs associados a um nome de domínio:

  ```bash
  ./goipsearch ip --host example.com
  ```

  - Para buscar nomes de domínio associados a um IP:

  ```bash
  ./goipsearch servidores --host 8.8.8.8
  ```

## Contribuindo

Contribuições são bem-vindas! Se você tiver sugestões de melhorias, abra uma issue para discutir suas ideias ou envie um pull request com suas alterações.

## Licença

Este projeto está licenciado sob a [MIT License](https://opensource.org/licenses/MIT). Sinta-se à vontade para usar, modificar e distribuir o código conforme necessário.
