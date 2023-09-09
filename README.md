sudo snap install hugo

cd docs
git submodule add https://github.com/theNewDynamic/gohugo-theme-ananke.git themes/ananke

curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash
nvm install v18.17.1

npm install -g autoprefixer
npm install -g postcss-cli
npm install -g postcss