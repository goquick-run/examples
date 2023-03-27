# Exemplos

### **Bem-vindo ao repositório de exemplos do Quick!**

Este repositório contém exemplos práticos de como utilizar a biblioteca Quick em Go, uma biblioteca de teste baseada em propriedades que permite escrever testes mais robustos e completos para sua aplicação.

Os exemplos apresentados aqui mostram como utilizar o Quick em diferentes tipos de testes, desde simples até mais complexos, ajudando a começar rapidamente e aprender as melhores práticas de teste.

O Quick é desenvolvido por **jeffotoni** e é uma excelente opção para escrever testes em Go, encontrando falhas em sua aplicação que podem não ser encontradas em testes tradicionais e aumentando a qualidade do código.

Sinta-se à vontade para explorar o repositório, contribuir com seus próprios exemplos e melhorias para a biblioteca Quick. Obrigado por usar Quick!

## Quais exemplos você encontrará no repositório?

  * [Group](/group/)
  * [Middleware Dinamico](/middleware.dinamico/)
  * [Middleware](/middleware/)
  * [Delete](quick.delete/)
  * [Get](quick.get/)
  * [Post](quick.post/)
  * [Put](quick.put/)
  * [Regex](quick.regex/)
  * [Start](quick.start/)

```go
package main

import "github.com/jeffotoni/quick"

func main() {
  app := quick.New()
  app.Get("/v1/user", func(c *quick.Ctx) error {
    c.Set("Content-Type", "application/json")
    return c.Status(200).SendString("Quick em ação com Cors❤️!")
  })
  app.Listen("0.0.0.0:8080")
}

```
