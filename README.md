# Description

repository going through postgres tutorial

# Task requirements

- [original documents with all tasks](documentation/PostgresqlTraining.kindOfEnglish.docx)

# Mkdocs

i document my going through in mkdocs, which is available if running
- python3 -m venv .venv
- poetry install
- mkdocs serve

once i finish completing it, i will publish results in github pages
Though... i don't document its going through in English

# Golang fakery data spamming

I complete the tutorial in Python and Golang languages. I started intially with python but discovered i can't complete performance measuring part of tutorial in a reliable way without generating fake data faster.

my ultra speed fakery data spamming code to postgres for sharing.
- [here is unit test that just validates data being created](<https://github.com/darklab8/darklab_training_postgresql/blob/68c7bec39e72ffd76e0e8a4ce7a43c9be1264f4f/golang/task2_test.go#L15>)
- [here is shared fixture that runs only one time for all tests (one time for testing session) to create temporal db and fill with fake data](<https://github.com/darklab8/darklab_training_postgresql/blob/68c7bec39e72ffd76e0e8a4ce7a43c9be1264f4f/golang/fixture_temp_db_test.go#L39>)
- [here is code defining which SQL tables to fill, and ties them in terms of foreign keys to each other](<https://github.com/darklab8/darklab_training_postgresql/blob/68c7bec39e72ffd76e0e8a4ce7a43c9be1264f4f/golang/fixtures.go#L60>)
- [here is algorithm itself of batch creating. Works for GORM and BUN orm frameworks as example in previous code shows](<https://github.com/darklab8/darklab_training_postgresql/blob/68c7bec39e72ffd76e0e8a4ce7a43c9be1264f4f/golang/shared/bulk_create.go#L75>)
- [it ultra spams those jobs in goroutine raised workers](<https://github.com/darklab8/darklab_training_postgresql/blob/68c7bec39e72ffd76e0e8a4ce7a43c9be1264f4f/golang/shared/bulk_create.go#L24>)

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
