bash

sudo apt-get update
sudo apt upgrade
sudo apt-get install zsh

sudo apt-get install curl
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

sudo apt-get install git
git clone https://github.com/denysdovhan/spaceship-prompt.git "$ZSH_CUSTOM/themes/spaceship-prompt"
ln -s "$ZSH_CUSTOM/themes/spaceship-prompt/spaceship.zsh-theme" "$ZSH_CUSTOM/themes/spaceship.zsh-theme"

echo "Set ZSH_THEME='spaceship'" >> ~/.zshrc
echo "" >> ~/.zshrc
echo "SPACESHIP_PROMPT_ORDER=(" >> ~/.zshrc
echo "  user          # Username section" >> ~/.zshrc
echo "  dir           # Current directory section" >> ~/.zshrc
echo "  host          # Hostname section" >> ~/.zshrc
echo "  git           # Git section (git_branch + git_status)" >> ~/.zshrc
echo "  hg            # Mercurial section (hg_branch  + hg_status)" >> ~/.zshrc
echo "  exec_time     # Execution time" >> ~/.zshrc
echo "  line_sep      # Line break" >> ~/.zshrc
echo "  vi_mode       # Vi-mode indicator" >> ~/.zshrc
echo "  jobs          # Background jobs indicator" >> ~/.zshrc
echo "  exit_code     # Exit code section" >> ~/.zshrc
echo "  char          # Prompt character" >> ~/.zshrc
echo ")" >> ~/.zshrc
echo "SPACESHIP_USER_SHOW=always" >> ~/.zshrc
echo "SPACESHIP_PROMPT_ADD_NEWLINE=false" >> ~/.zshrc
echo "SPACESHIP_CHAR_SYMBOL="❯"" >> ~/.zshrc
echo "SPACESHIP_CHAR_SUFFIX=" "" >> ~/.zshrc

zsh
upgrade_oh_my_zsh