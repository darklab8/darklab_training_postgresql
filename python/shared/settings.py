from python.utils.config_parser import ConfigParser
from python.utils.logger import Logger

logger = Logger(console_level="DEBUG", name=__name__)
config = ConfigParser(settings_prefix="training")

DATABASE_USER = config.get("database_username","postgres")
DATABASE_PASSWORD = config.get("database_password", "postgres")
DATABASE_HOST = config.get("database_host", "localhost")
DATABASE_URL = config.get(
    "database_url", f"{DATABASE_USER}:{DATABASE_PASSWORD}@{DATABASE_HOST}/"
)
DATABASE_DEBUG = bool(config.get("database_debug", ""))