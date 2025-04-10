import psycopg
from psycopg_pool import AsyncConnectionPool

class PGPool:
    _instance = None

    def __new__(cls, connString: str = None):
        if cls._instance is None:
            if connString is None:
                raise ValueError("First call must include a DSN")
            cls._instance = super().__new__(cls)
            cls._instance._pool = AsyncConnectionPool(connString)
        return cls._instance

    def connect(self) -> psycopg.connection:
        return self._pool.getconn()

    def close(self):
        self._pool.close()