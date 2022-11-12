#!/usr/bin/env python3.10
from enum import Enum, auto
import subprocess
import argparse
import logging
import os
from types import SimpleNamespace

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
    dev_test = "testing (for postgresql) (TODO add to CI)"
    shell_docs_build = "build (for documentation)"
    shell_docs_dev = "rendering in dev env"
    shell_poetry_export = "regenerate requirements.txt for docker image"

def shell(cmd):
    subprocess.run(cmd, shell=True, check=True)

def main():
    parser = argparse.ArgumentParser()

    actions_parser = parser.add_subparsers(dest="action", required=True,)
    actions = SimpleNamespace()
    for action in Action:
        setattr(actions, action.name, actions_parser.add_parser(action.name, help=action.value))
    
    args = parser.parse_args()
    logger.info(f"{args=}")

    match Action[args.action]:
        case Action.dev_shell:
            shell("docker-compose build -- app && docker-compose run --rm app sh ; docker-compose down")
        case Action.dev_test:
            shell("docker-compose build -- app && docker-compose run --rm app pytest ; docker-compose down")
        case Action.dev_down:
            shell("docker-compose down")
        case Action.shell_docs_build:
            shell("mkdocs build")
        case Action.shell_docs_dev:
            shell("mkdocs serve")
        case Action.shell_poetry_export:
            shell("poetry export --without-hashes --format=requirements.txt > requirements.txt")


if __name__=="__main__":
    main()

