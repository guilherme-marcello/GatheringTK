package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "strings"
)

const(
        blue="\033[36m"
        white="\033[97;1m"
        cyan="\033[96;1m"
        red = "\033[91;1m"
        green = "\033[92;1m"
        yellow = "\033[93;1m"
        clear = "\033[H\033[2J"
        menu = "\033[36m"
        reset = "\033[0m"
        option = "\033[33m"
)

func print(v string){
    fmt.Print(v);}

func help_menu(){
    fmt.Printf("\n %s \033[4m MIXER HELP\n%s", cyan, reset)
    fmt.Printf("%s*************************************************************************\n", menu)
    fmt.Printf("** %s'init' | 'file') %sCriar o ficheiro a ser utilizado. ('init filename')\n", option, menu)
    fmt.Printf("** %s'run'|'--r') %sExecutar o programa. ('run string1 string2 string3 ...')\n", option, menu)
    fmt.Printf("** %s'clear'|'-c'|'limpar') %sLimpar ecrã. ('clear')\n", option, menu)
    fmt.Printf("** %s'back'|'voltar'|'exit'|'-e'|'sair'|':q') %sVoltar atrás. ('back')\n", option, menu) 
    fmt.Printf("** %s'help'|'-h'|'--help'|'ajuda'|'?') %sExibir a ajuda. ('help')\n", option, menu)
    fmt.Printf("*************************************************************************\n")
}

func noContent(text string, command string) bool{
    lenCommand := len(command)
    textCommand := text[lenCommand:len(text)]
    if len(strings.TrimSpace(textCommand)) == 0 {
        fmt.Println(red+"[!] ",command,": falta de argumentos após o comando.")
        return true
    }
    return false   
}

func file_not_in_dir(dirname string, picked string) bool{
    var status bool
    status = true
    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if picked == file.Name() {
            status = false
        }
    }
    return status
}	


func main() {
  reader := bufio.NewReader(os.Stdin)
  var P bool
  var f *os.File
  var ficheiro string
  for {
    print(white)
    print("gatheringtk("+blue+"mixer"+white+") >> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1) 
    nomes := strings.Split(text," ")
    if (strings.Compare("", text) != 0) {
        switch nomes[0] {
        case "init","file":
            if noContent(text,"init") {
                break
            }
            ficheiro = nomes[1]+".txt"
            if file_not_in_dir("mixer/OUTPUT/",ficheiro){
                fi, err := os.Create("mixer/OUTPUT/"+ficheiro)
                f = fi
                if err != nil {
                    fmt.Println(err)
                    return
                    }
                defer f.Close()
                P = true
                wd, err := os.Getwd()
                if err != nil {
                    log.Println(err)
                }
                fmt.Println(green+"[+] Ficheiro "+ficheiro+" criado e gravado em "+wd+"/mixer/OUTPUT."+white)            
            } else {fmt.Println(yellow+"[*] Existe um ficheiro com o nome "+ficheiro+" em mixer/OUTPUT.\n"+red+"[!] Inicialização não realizada."+white)}
        case "help", "-h", "--help", "ajuda", "?":
            help_menu()
        case "back","voltar","exit", "-e", "sair", ":q":
            fmt.Println(yellow+"[*] O mixer será encerrado.")
            fmt.Println(green+"[+] Os desenvolvedores agradecem a sua utilização!")
            os.Exit(0)
        case "clear","-c","limpar":
            print(clear)
            break
        case "run", "--r":
            if P != true {
                fmt.Println(red+"[!] É preciso selecionar o nome do ficheiro.")
                break
            }
            if noContent(text,"run") {
                break
            }
            nomes = nomes[1:len(nomes)]
            for idxG,nomeGeral := range nomes{
                for idxInd,nome := range nomes{
                    if idxG != idxInd{        
                        n, err := f.WriteString(nomeGeral+nome+"\n"+strings.ToUpper(nomeGeral)+strings.ToUpper(nome)+"\n"+strings.ToLower(nomeGeral)+strings.ToLower(nome)+"\n"+strings.ToLower(nomeGeral)+"."+strings.ToLower(nome)+"\n"+strings.ToLower(nomeGeral)+"-"+strings.ToLower(nome)+"\n"+ strings.ToUpper(nomeGeral)+"."+strings.ToUpper(nome)+"\n"+strings.ToUpper(nomeGeral)+"-"+strings.ToUpper(nome)+"\n")
                        if err != nil {
                            fmt.Println(err)
                            log.Fatal(red+"[!!!] Impossível escrever no ficheiro.")
                            return
                        }
                        fmt.Println(yellow+"[*] ",n," bytes adicionados ao ficheiro."+white)
                   } 
                }
            }
            fmt.Println(green+"[+] Todos os 'usernames' possíveis adicionados ao ficheiro ",ficheiro,".")
        default:
            fmt.Println(red+"[!] Comando inexistente. Consulte os comandos com o comando help.")
        }

  }

 }

}



