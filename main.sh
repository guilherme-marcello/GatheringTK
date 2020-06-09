#!/bin/sh

cyan=`echo "\033[96;1m"`
yellow=`echo "\033[93;1m"`
blue_menu=`echo "\033[36m"`
white=`echo "\033[97;1m"`
normal=`echo "\033[m"`
menu=`echo "\033[36m"`
number=`echo "\033[33m"` 
sred=`echo "\033[01;31m"`


banner(){
echo "  ▄▀  ██     ▄▄▄▄▀ ▄  █ ▄███▄   █▄▄▄▄ ▄█    ▄     ▄▀         ▄▄▄▄▀ █  █▀ ";
echo "▄▀    █ █ ▀▀▀ █   █   █ █▀   ▀  █  ▄▀ ██     █  ▄▀        ▀▀▀ █    █▄█   ";
echo "█ ▀▄  █▄▄█    █   ██▀▀█ ██▄▄    █▀▀▌  ██ ██   █ █ ▀▄          █    █▀▄   ";
echo "█   █ █  █   █    █   █ █▄   ▄▀ █  █  ▐█ █ █  █ █   █        █     █  █  ";
echo " ███     █  ▀        █  ▀███▀     █    ▐ █  █ █  ███        ▀        █   ";
echo "        █           ▀            ▀       █   ██                     ▀    ";
echo "       ▀                                                                 ";
echo "\n                    developer: Guilherme Marcello";
}

help_menu(){
    printf "${cyan} \033[4m HELP "
    printf "\n${menu}*********************************************${normal}\n"
    printf "${menu}**${number} 'clear'|'-c'|'limpar')${menu} Limpar ${normal}\n"
    printf "${menu}**${number} 'help'|'-h'|'ajuda')${menu} Ajuda ${normal}\n"
    printf "${menu}**${number} 'show options'|'options')${menu} Exibir as opções ${normal}\n"
    printf "${menu}**${number} 'use')${menu} Selecionar ferramenta ${normal}\n"
    printf "${menu}**${number} 'exit'|'sair'|'-s')${menu} Sair do programa ${normal}\n"
    printf "${menu}*********************************************${normal}\n\n"
}

show(){
$1
}

main_menu(){
    printf "\n${menu}*********************************************${normal}\n"
    printf "${menu}**${number} 1)${menu} Mixer ${normal}\n"
    printf "${menu}**${number} 99)${menu} Sair ${normal}\n"
    printf "${menu}*********************************************${normal}\n"
    printf "Escolha uma opção. ${normal}\n"

}

openTerminal(){
read -p "${white}gatheringtk(${blue_menu}$3${white}) >> " choice
      case $choice in
        "clear"|"-c"|"limpar") clear;
        ;;
        "help"|"-h"|"ajuda") show help_menu; 
        ;;
        "show options"|"options") show $2; continue;
        ;;
        1|'use 1') $1 
        ;;
        99|'exit'|'sair'|'-s'|'use 99') echo "${sred}[!] O programa será encerrado.";exit;
        ;;
        '') continue
        ;;
      esac
}

execMixer(){
go run mixer/main.go
}

main(){
show banner	
while [ "$choice"!='' ]
    do
         openTerminal execMixer main_menu /	
    done 
}

main
