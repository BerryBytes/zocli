#!/bin/bash
set -e
set -eo pipefail
print_error_message() {
    echo "An error occurred. Please visit https://docs.01cloud.io/services/cli/ for assistance."
}

trap 'if [ $? -ne 0 ]; then print_error_message; fi' EXIT

zoclipath="$HOME/.01cloud/"
printf "\\e[1mINSTALLATION\\e[0m\\n"
if [ -d $zoclipath ]; then
  echo "Previous installation detected."
else
  mkdir ${zoclipath}
  echo "Created " $zoclipath " folder."
fi

printf "Downloading zocli."
print_dots() {
  while true; do
    printf "."
    sleep 1
  done
}
print_dots &

ARCH=$(uname -m)
if [ "$ARCH" = "arm64" ]; then
  ARCH="arm64"
elif [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

OS_NAME=$(uname -s)
VERSION=${1:-latest}  # Accept version as an argument, default to 'latest' if not provided
if [ "$OS_NAME" = "Darwin" ]; then
  if curl -sS -L "https://github.com/berrybytes/zocli/releases/${VERSION}/download/zocli_Darwin_${ARCH}" -o $HOME/.01cloud/zocli; then
    kill "$!" 2>/dev/null
    echo ""
    echo "Download completed"
  else
    kill "$!" 2>/dev/null
    echo "Downloading failed. Check your internet connection."
    exit 1
  fi
elif [ "$OS_NAME" = "Linux" ]; then
  if curl -sS -L "https://github.com/berrybytes/zocli/releases/${VERSION}/download/zocli_Linux_${ARCH}" -o $HOME/.01cloud/zocli; then
    kill "$!" 2>/dev/null
    echo ""
    echo "Download completed"
  else
    kill "$!" 2>/dev/null
    echo "Downloading failed. Check your internet connection."
    exit 1
  fi
elif ["$OS_NAME" = "Windows"]; then
  kill "$!" 2>/dev/null
  echo "Direct installer for windows is not yet supported: Check GITHUB RELEASES for manual installation"
  echo "RELEASES: https://github.com/BerryBytes/zocli/releases"
  exit 1
else
    kill "$!" 2>/dev/null
    echo "Unsupported operating system: $OS_NAME"
    exit 1
fi

chmod +x $HOME/.01cloud/zocli
if [ $? -eq 0 ]; then
  echo ""
else
    echo "Error: chmod failed."
    exit 1
fi

CURRENT_SHELL="$SHELL"

if [[ "$CURRENT_SHELL" = "/bin/bash" || "$CURRENT_SHELL" = "/usr/bin/bash" ]]; then
      echo "Detected Bash shell."
  CONFIG_FILE="$HOME/.bashrc"
    if grep -q ".01cloud" "$CONFIG_FILE"; then
      echo "The PATH is already set in $CONFIG_FILE."
    else
      echo "export PATH=\"$HOME/.01cloud:$PATH\"" >> $CONFIG_FILE
      source $CONFIG_FILE
    fi
elif [[ "$CURRENT_SHELL" = "/bin/zsh" || "$CURRENT_SHELL" = "/usr/bin/zsh" ]]; then
      echo "Detected Zsh shell."
  CONFIG_FILE="$HOME/.zshrc"
    if grep -q ".01cloud" "$CONFIG_FILE"; then
      echo "The PATH is already set in $CONFIG_FILE."
    else
      echo "export PATH=\"$HOME/.01cloud:$PATH\"" >> $CONFIG_FILE
  echo ""
      zsh
    fi
elif [[ "$CURRENT_SHELL" = "/bin/fish" || "$CURRENT_SHELL" = "/usr/bin/fish" ]]; then
    echo "Detected Fish shell."
    set -U fish_user_paths \"$HOME/.01cloud\" $fish_user_paths
else
  printf "\\e[1mFAILURE\\e[0m\\n"
  echo "Unsupported shell detected: $CURRENT_SHELL"
  echo "Please set the PATH manually."
  echo "See https://docs.01cloud.io/services/cli/"
fi

echo ""
printf "\\e[1m------INSTALLATION COMPLETED------\\e[0m\\n"
echo ""
printf "\\e[1mSUMMARY\\e[0m\\n"
echo "    zocli is an official CLI tool for managing 01Cloud resources."
echo ""
printf "\\e[1mUSAGE\\e[0m\\n"
echo "    zocli --help"
echo ""
printf "\\e[1mUNINSTALL\\e[0m\\n"
echo "    everything is installed into ~/.01cloud/,"
echo "    so you can remove it like so:"
echo ""
echo "    rm -rf ~/.01cloud/"
echo ""
printf "\\e[1mGETTING STARTED\\e[0m\\n"
printf "    See \\e[34mhttps://docs.01cloud.io/services/cli/\\e[0m\\n"
echo ""
printf "\\e[1mTIP\\e[0m\\n"
printf "    Inorder to use zocli in this terminal, please run \\e[34msource ~/.bashrc\\e[0m or you own shell CONFIG_FILE\n"
echo ""
exit
