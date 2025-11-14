# 🚀 Projeto de Onboarding: Git & Golang

## 👋 Boas-vindas ao Time!

Parabéns por se juntar à equipe! Estamos muito empolgados em ter você conosco.

Este repositório é seu primeiro projeto de integração e serve a dois propósitos principais:

1.  **Praticar nosso fluxo de trabalho com Git:** Você vai clonar, criar uma *branch*, fazer *commits*, dar *push* e abrir um *Pull Request* (PR).
2.  **Aquecer os motores com Go:** Você vai criar um programa "Olá, Mundo!" simples e depois completar alguns desafios básicos da linguagem.

Não se preocupe se algo parecer novo; estamos aqui para ajudar. O objetivo é aprender!

---

## 🎯 Parte 1: O Primeiro Pull Request

Sua primeira tarefa é **adicionar uma nova pasta neste repositório com o seu nome de usuário e, dentro dela, criar um programa básico em Go.**

### 1. A Tarefa em Go (Golang)

Dentro da sua pasta, crie um arquivo chamado `main.go`. Este programa deve:

* Estar no pacote `main`.
* Ter uma função `main`.
* Quando executado (com `go run main.go`), ele deve imprimir no console uma saudação personalizada.

**Exemplo de `main.go`:**

```go
package main

import "fmt"

func main() {
    // Sinta-se à vontade para personalizar esta mensagem!
    fmt.Println("Olá, equipe! Meu nome é [Seu Nome] e estou pronto(a) para começar!")
}
```

### 2. O Fluxo de Trabalho no Git (Passo a Passo)

Siga estas etapas para concluir o desafio.

**Passo 1: Clonar o Repositório**
Clone este projeto para a sua máquina local. (Use SSH ou HTTPS).

```bash
git clone [URL_DO_SEU_REPOSITORIO_AQUI]
```

**Passo 2: Criar sua Branch**
**Importante:** Nunca trabalhe diretamente na *branch* `main`. Sempre crie uma *branch* nova para suas alterações.

Use um nome descritivo para sua *branch*, por exemplo: `feature/add-nome-sobrenome`.

```bash
# Vá para a pasta do projeto
cd [nome-do-repositorio]

# Crie e mude para sua nova branch
git checkout -b feature/add-seu-nome-aqui
```

**Passo 3: Criar sua Pasta e Arquivo**
Agora, crie sua pasta pessoal e o arquivo `main.go` dentro dela, conforme descrito na "Tarefa em Go".

**Passo 4: Adicionar e "Comitar" suas Mudanças**
Adicione seus arquivos ao "stage" do Git e, em seguida, crie um *commit* (um "instantâneo" das suas mudanças) com uma mensagem clara.

```bash
# Adiciona sua pasta (substitua pelo nome da sua pasta)
git add sua-pasta-aqui/

# Cria o commit
git commit -m "feat: Adiciona o programa inicial de [Seu Nome]"
```

**Passo 5: Enviar suas Mudanças (Push)**
Envie sua *branch* (e seus *commits*) para o repositório remoto (GitHub/GitLab/etc.).

```bash
git push origin feature/add-seu-nome-aqui
```

**Passo 6: Abrir um Pull Request (PR)**
Vá para a interface do GitHub (ou similar). Você verá um aviso sugerindo a criação de um **Pull Request** a partir da *branch* que você acabou de enviar.

* Abra o Pull Request.
* Coloque um título claro (ex: "Adiciona programa Go de [Seu Nome]").
* No corpo do PR, marque seu mentor ou o líder técnico para revisão.

---

## 💻 Parte 2: Desafios de Golang

Ótimo! Seu primeiro PR está aberto e seu `main.go` funciona. Agora, vamos expandir seu conhecimento em Go.

**Modifique o seu arquivo `main.go`** para completar os desafios abaixo.

> **Importante:** Para cada desafio que você completar, crie um **novo commit** na sua *branch* e dê `push`. Isso atualizará seu Pull Request automaticamente e nos permite ver seu progresso.
>
> ```bash
> # Após completar o Desafio 1:
> git add sua-pasta-aqui/main.go
> git commit -m "feat: Adiciona função de soma"
> git push
> ```

### Desafio 1: Função de Soma

1.  Abaixo da sua função `main`, crie uma nova função chamada `soma` que recebe dois parâmetros `int` (ex: `a` e `b`) e retorna um `int`.
2.  A função deve retornar a soma dos dois números.
3.  Na sua função `main`, chame a função `soma` (ex: `soma(5, 5)`) e imprima o resultado no console.
    * *Saída esperada: "A soma é: 10"*

### Desafio 2: Imprimir um Slice (Array)

1.  Na sua função `main`, declare um *slice* de *strings* com pelo menos 3 nomes de frutas.
    * Ex: `frutas := []string{"Maçã", "Banana", "Laranja"}`
2.  Use um loop `for...range` para iterar sobre o *slice*.
3.  Imprima cada item do *slice* no console, um por linha.

### Desafio 3: Imprimir um Triângulo

1.  Na sua função `main`, use um loop `for` (ou loops aninhados) para imprimir o seguinte padrão de asteriscos no console:
    ```
    *
    **
    ***
    ****
    *****
    ```

### Desafio Bônus: Trabalhando com Mapas

1.  Crie um `map` onde a chave é uma `string` (nome de um colega) e o valor é um `int` (idade).
2.  Adicione alguns colegas a este mapa.
3.  Use um loop `for...range` para imprimir o nome e a idade de cada pessoa no mapa.

---

## 🏋️‍♀️ Parte 3: Praticando o Fluxo de Trabalho (Git)

Com seu PR aberto e vários *commits* adicionados, vamos praticar algumas tarefas essenciais de Git.

### Desafio 1: Sincronizando com `merge`

É comum que a *branch* `main` receba atualizações (de outros colegas) enquanto você ainda está trabalhando na sua. Você precisa trazer essas atualizações para sua *branch* antes de finalizar seu PR.

1.  **Sincronize sua `main` local:** Primeiro, garanta que sua *branch* local `main` está atualizada.
    ```bash
    git checkout main
    git pull origin main
    ```
2.  **Volte para sua *branch*:**
    ```bash
    git checkout feature/add-seu-nome-aqui
    ```
3.  **Traga as atualizações (Merge):** Agora, traga as atualizações da `main` para a sua *branch* atual.
    ```bash
    git merge main
    ```
4.  **Conflitos:** Se o Git acusar um `CONFLICT` (conflito), isso significa que você e outra pessoa alteraram as mesmas linhas de um mesmo arquivo.
    * Abra os arquivos com conflito (o Git informará quais são).
    * Procure as marcações (`<<<<<<<`, `=======`, `>>>>>>>`).
    * Edite o arquivo para manter o código correto e **remova as marcações**.
    * Salve o arquivo e continue o *merge*:
    ```bash
    git add .
    git merge --continue
    ```

### Desafio 2: Corrigindo o Último Commit (`--amend`)

Você fez um `commit`, mas percebeu que digitou a mensagem errada ou esqueceu de adicionar um arquivo? Use o `amend`.

1.  **Faça sua correção:** Se esqueceu um arquivo, adicione-o:
    ```bash
    git add arquivo-que-esqueci.go
    ```
2.  **Use `amend`:** Este comando "emenda" (corrige) o *commit* anterior.

    ```bash
    # O '--amend' vai abrir o editor de texto para você corrigir a mensagem.
    git commit --amend
    
    # Se você só queria adicionar arquivos e a mensagem estava OK:
    git commit --amend --no-edit
    ```
4.  **Atenção:** Como você já deu `push` nesse *commit*, o `amend` reescreve o histórico. Você precisará "forçar" o *push*.
    ```bash
    # Force-push com segurança (só atualiza se ninguém mais mexeu na branch)
    git push --force-with-lease
    ```

### Desafio 3: Sincronizando com `rebase`

O `git merge` (Desafio 1) cria um "commit de merge". Muitos times preferem o `git rebase` para manter um histórico linear. O `rebase` "pega" os *commits* da sua *branch* e os reaplica *no topo* da *branch* `main`.

1.  **Busque as atualizações:** `git fetch origin`
2.  **Inicie o Rebase:** Em vez de `merge`, use `rebase`.
    ```bash
    # (Estando na sua branch)
    # Isso significa: "Reaplique meus commits no topo da 'origin/main'"
    git rebase origin/main
    ```
3.  **Conflitos durante o `rebase`:** O `rebase` aplica seus *commits* um por um. Se houver um conflito, ele vai parar.
    * Resolva o conflito (como no Desafio 1).
    * Use `git add .` para marcar como resolvido.
    * Use `git rebase --continue` para pular para o próximo *commit*.
    * (Se ficar confuso, você sempre pode parar com `git rebase --abort`).

---

## 📁 Estrutura Esperada do Repositório

Após alguns *commits*, o repositório deverá ficar parecido com isto:

```
/
├── .gitignore
├── README.md
│
├── ana.silva/
│   └── main.go
│
├── bruno.costa/
│   └── main.go
│
└── (sua-pasta-aqui)/
    └── main.go
```

## 📚 Recursos Úteis

* [A Tour of Go](https://go.dev/tour/welcome/1) (Oficial - Excelente para começar)
* [Pro Git Book](https://git-scm.com/book/pt-br/v2) (Livro completo sobre Git, em português)
* [Git Cheat Sheet](https://education.github.com/git-cheat-sheet-education.pdf) (Folha de consulta rápida do Git)
* [Learn Git Branching](https://learngitbranching.js.org/?locale=pt_BR) (Visual e interativo)

Dúvidas? Pergunte no chat do time! Boa sorte!
