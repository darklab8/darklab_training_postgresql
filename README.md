# Description

repository going through postgres tutorial

# Python requirements

- python3.10
- poetry https://python-poetry.org/docs/#installing-with-the-official-installer (can be replaced with standard pip as long as not changed deps)
- docker-compose (only for pytest tests + postgresql code)

# Golang requirements

# Hugo requirements

sudo snap install hugo

cd docs
git submodule add https://github.com/theNewDynamic/gohugo-theme-ananke.git themes/ananke

curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash
nvm install v18.17.1

npm install -g autoprefixer
npm install -g postcss-cli
npm install -g postcss
