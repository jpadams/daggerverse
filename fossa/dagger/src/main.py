"""Base Container and analysis function for fossa-cli

Today this module has a base() function that returns a Container and an
analyze() function that can take a token for the Fossa cloud service to
update analysis or just print the information to stdout locally. You must
provide a directory of source code for the tool to analyze.

In the future, the module may add functions to take full advantage of fossa-cli
capabilities including scanning container imgages.

Note: as of this writing, fossa-cli only supports amd64 on linux and not arm so
we explicitly specify that the container that we load the cli into should be
linux/amd64.
"""

import time
import dagger
from dagger import dag, function, object_type


@object_type
class Fossa:
    @function
    async def analyze(self, source: dagger.Directory, fossa_token: dagger.Secret | None) -> str:
        """Analyzes a directory of source code, prints to stdout, optionally uploads results to Fossa"""
        ctr = self.base()
        ctr = ctr.with_mounted_directory("/src", source).with_workdir("/src")

        cmd = ["fossa", "analyze"]

        if fossa_token is None:
           cmd.append("--output") 
        else:
           ctr = ctr.with_env_variable("FOSSA_API_KEY", await fossa_token.plaintext())    
        return await (
            ctr.with_env_variable("CACHEBUSTER", str(time.time())) 
            .with_exec(cmd)
            .stdout()
        )

    @function
    def base(self) -> dagger.Container:
        """Returns a minimal Container with fossa-cli"""
        return (
            dag.container(platform=dagger.Platform("linux/amd64"))
            .from_("alpine:latest", )
            .with_exec(["apk", "add", "curl", "bash"])
            .with_exec(["sh", "-c", "curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/fossas/fossa-cli/master/install-latest.sh | bash"])
        ) 
