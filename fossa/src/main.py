import time
import dagger
from dagger.mod import function

@function
async def analyze(fossa_token: dagger.Secret, source = dagger.Directory) -> str:
    # Example usage: "dagger call analyze --fossa-token <FOSSA token> --source <directory path>"
    ctr = base()
    return await (
        ctr.with_env_variable("FOSSA_API_KEY", await fossa_token.plaintext())	
        .with_mounted_directory("/src", source)
        .with_workdir("/src")
        .with_env_variable("CACHEBUSTER", str(time.time()))
        .with_exec(["fossa", "analyze"])
        .stdout()
    )

@function
def base() -> dagger.Container:
    return (
        dagger.container(platform=dagger.Platform("linux/amd64"))
        .from_("alpine:latest", )
        .with_exec(["apk", "add", "curl", "bash"])
        .with_exec(["sh", "-c", "curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/fossas/fossa-cli/master/install-latest.sh | bash"])
    ) 
