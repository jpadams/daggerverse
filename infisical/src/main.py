import os
import dagger
from dagger import dag, function
from infisical_client import ClientSettings, InfisicalClient, GetSecretOptions

@function
async def get_secret(name: str, token: dagger.Secret, project_id: str, env: str, path: str = "") -> dagger.Secret:
    """Get a secret from an Infisical project using secret name, project token, env, and path"""
    inf_client = InfisicalClient(ClientSettings(
        access_token=await token.plaintext(),
    ))
    secret = inf_client.getSecret(options=GetSecretOptions(
        environment=env,
        project_id=project_id,
        secret_name=name,
        path=path
    ))
    return dag.set_secret("val", secret.secret_value)

@function
async def test(token: str, project_id: str) -> str:
    """Insecure test using default Infisical project and plaintext token"""
    return await dagger.Secret.plaintext(
        await get_secret(
            name="DATABASE_URL",
            token=dag.set_secret("tok", token),
            project_id=project_id,
            env="dev",
            path="/")
         )
