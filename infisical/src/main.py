import os
import dagger
from dagger.mod import function
from infisical import InfisicalClient

@function
async def get_secret(name: str, token: dagger.Secret, env: str, path: str) -> dagger.Secret:
    """Get a secret from an Infisical project using secret name, project token, env, and path"""
    inf_client = InfisicalClient(token=await dagger.Secret.plaintext(token))
    return dagger.set_secret("val", inf_client.get_secret(name, environment=env, path=path).secret_value)

@function
async def test(token: str) -> str:
    """Insecure test using default Infisical project and plaintext token"""
    return await dagger.Secret.plaintext(
        await get_secret(
            name="DATABASE_URL",
            token=dagger.set_secret("tok", token),
             env="dev",
             path="/")
         )
