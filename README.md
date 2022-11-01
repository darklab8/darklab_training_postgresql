# requirements

- python3.10
- poetry https://python-poetry.org/docs/#installing-with-the-official-installer (can be replaced with standard pip as long as not changed deps)
- docker-compose (only for pytest tests + postgresql code)
# dev environment (for documentation)

- `poetry install` (or p`ip install -r requirements.txt`, if u aren't planning add new python deps)
- `poetry shell` (only if installed via poetry)
- `mkdocs serve`

# build (for documentation)

- `pip install -r requirements.txt`
- `mkdocs build`
- (read .gitlab-ci.yml for actual relevant code that does it in CI)

# dev env (for postgresql)

- `docker-compose build -- app && docker-compose run -v $(pwd):/code -u 0 --rm app sh && docker-compose down`

# testing (for postgresql) (TODO add to CI)

- `docker-compose build -- app && docker-compose run -v $(pwd):/code -u 0 --rm app pytest && docker-compose down`