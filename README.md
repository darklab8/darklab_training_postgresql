# requirements

- python3.10
- poetry https://python-poetry.org/docs/#installing-with-the-official-installer

# dev environment

- poetry install (or pip install -r requirements.txt, if u aren't planning add new python deps)
- poetry shell (only if installed via poetry)
- mkdocs serve

# build

- pip install -r requirements
- mkdocs build
- (read .gitlab-ci.yml for actual relevant code that does it in CI)