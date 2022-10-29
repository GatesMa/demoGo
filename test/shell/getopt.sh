#!/bin/bash

func() {
    echo "Usage:"
    echo "test.sh [-j S_DIR] [-m D_DIR]"
    echo "Description:"
    echo "S_DIR,the path of source."
    echo "D_DIR,the path of destination."
    exit -1
}

#upload="false"

while getopts 'h:j:m:u' OPT; do
    case $OPT in
        j) S_DIR="$OPTARG";;
        m) D_DIR="$OPTARG";;
        u) upload="true";;
        h) func;;
        ?) func;;
    esac
done

echo $S_DIR
echo $D_DIR
echo $upload


###################
echo $SHELL
if [[ "$SHELL" == "/bin/bash" ]]; then
    # 绝对路径
    export RS_ROOT=$(cd $(dirname "${BASH_SOURCE}");pwd)
#    if ! grep -q "source ${RS_ROOT}/rsrc" ~/.bashrc; then
#        echo "source ${RS_ROOT}/rsrc" >> ~/.bashrc
#        echo "" >> ~/.bashrc
#    fi
#    if ! grep -q "source ${RS_ROOT}/rsrc" ~/.bash_profile; then
#        echo "source ${RS_ROOT}/rsrc" >> ~/.bash_profile
#        echo "" >> ~/.bash_profile
#    fi
elif [[ "$SHELL" == "/bin/zsh" ]]; then
    # https://stackoverflow.com/questions/9901210/bash-source0-equivalent-in-zsh
    # 绝对路径
    export RS_ROOT=$(cd $(dirname ${(%):-%N});pwd)
    if ! grep -q "source ${RS_ROOT}/rsrc" ~/.zshrc; then
        echo "source ${RS_ROOT}/rsrc" >> ~/.zshrc
        echo "" >> ~/.zshrc
    fi
else
    echo "unsupported shell: $SHELL"
fi

echo $RS_ROOT