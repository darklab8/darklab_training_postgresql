import pytest
from . import models

pytestmark = pytest.mark.django_db

def test_check_pytest():
    assert True

def test_creating_object():
    models.Example.objects.create(name="123")

    found_object = models.Example.objects.filter(name="123").first()

    assert found_object.name == "123"

    assert models.Example.objects.filter(name="1234").first() is None