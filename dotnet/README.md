# Usage

`ln -s /usr/local/share/dotnet/dotnet /usr/local/bin/`

`dotnet new -i Google.Cloud.Functions.Templates`

`dotnet new gcf-http`

or
```
mkdir HelloFunctions
cd HelloFunctions
dotnet new gcf-http
```

`dotnet run`

gcloud functions deploy dotnet-function --runtime dotnet3 --trigger-http --allow-unauthenticated --entry-point dotnet.Function