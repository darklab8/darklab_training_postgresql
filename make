#!/usr/bin/env python3.10
from enum import Enum, auto
import subprocess
import argparse
import logging
import os
from types import SimpleNamespace
import secrets

def get_logger():
    debug_level = os.environ.get("LOGGING_LEVEL","DEBUG")

    class LoggerLevels(Enum):
        DEBUG = logging.DEBUG
        INFO = logging.INFO 
        WARN = logging.WARNING 
        ERROR = logging.ERROR
        FATAL = logging.CRITICAL

    logger = logging.getLogger().getChild(__name__)
    logger.setLevel(LoggerLevels[debug_level].value)
    ch = logging.StreamHandler()
    formatter = logging.Formatter(" - ".join(["time:%(asctime)s","level:%(levelname)s","name:%(name)s","msg:%(message)s"]))
    ch.setFormatter(formatter)
    ch.setLevel(LoggerLevels[debug_level].value)
    logger.addHandler(ch)
    return logger

logger = get_logger()

class Action(Enum):
    dev_shell = "dev env (for postgresql)"
    dev_down = auto()
    ci_test = auto()
    dev_docs_build = "build (for documentation)"
    dev_docs_run = "rendering in dev env"
    shell_poetry_export = "regenerate requirements.txt for docker image"

def shell(cmd):
    subprocess.run(cmd, shell=True, check=True)

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--id", default=secrets.token_hex(2))

    actions_parser = parser.add_subparsers(dest="action", required=True,)
    actions = SimpleNamespace()
    for action in Action:
        setattr(actions, action.name, actions_parser.add_parser(action.name, help=action.value))
    
    args = parser.parse_args()
    logger.info(f"{args=}")

    match Action[args.action]:
        case Action.dev_shell:
            shell(f"docker-compose -p {args.id} build -- app && docker-compose -p {args.id} run --service-ports --rm -v $(pwd):/code app_with_pgadmin bash ; docker-compose -p {args.id} down")
        case Action.dev_down:
            shell("docker-compose down")
        case Action.dev_docs_build:
            shell(f"docker-compose -p {args.id} build -- app && docker-compose -p {args.id} run --rm -v $(pwd):/code app mkdocs build ; docker-compose -p {args.id} down")
        case Action.dev_docs_run:
            shell(f"docker-compose -p {args.id} build -- app && docker-compose -p {args.id} run -p 8000:8000 --rm -v $(pwd):/code app mkdocs serve -a 0.0.0.0:8000 ; docker-compose -p {args.id} down")
        case Action.shell_poetry_export:
            shell("poetry export --without-hashes --format=requirements.txt > requirements.txt")
        case Action.ci_test:
            try:
                shell(f"docker-compose -p {args.id} build -- app")
                shell(f"docker-compose -p {args.id} pull")
                shell(f"docker-compose -p {args.id} up -d")
                shell(f"docker-compose -p {args.id} run app ./wait_for_it.sh db:5432 -t 30 -- pytest")
            finally:
                shell(f"docker-compose -p {args.id} down")


if __name__=="__main__":
    main()

