# Regenerate the README screenshots.
#
# Uses headless Chrome against a running dev stack, so the images in the README
# are always reproducible from the actual application rather than hand-cropped
# and quietly going stale.
#
# Requires: API on :8080, web on :5173, database seeded with dev/seed.sql.
#
#   pwsh dev/screenshots.ps1                    # all shots, light theme
#   pwsh dev/screenshots.ps1 -Theme dark        # dark theme
#   pwsh dev/screenshots.ps1 -Only ask,course   # a subset
#
# Theme is forced through Chrome's preferred-color-scheme rather than by
# clicking the in-app toggle: the app reads the OS preference on first paint,
# so setting it at the browser level gives the same result without scripting
# the UI.

param(
  [ValidateSet("light", "dark")]
  [string]$Theme = "light",
  [string[]]$Only = @()
)

$ErrorActionPreference = "Stop"

$chrome = "C:\Program Files\Google\Chrome\Application\chrome.exe"
$out    = Join-Path $PSScriptRoot "..\docs\screenshots"
$base   = "http://localhost:5173"
$course = "22222222-2222-2222-2222-222222222222"

New-Item -ItemType Directory -Force -Path $out | Out-Null

$shots = @(
  @{ name = "dashboard";  url = "$base/" },
  @{ name = "course";     url = "$base/courses/$course" },
  @{ name = "questions";  url = "$base/courses/$course/questions?chapter=3" },
  # A real question, so the shot shows the feature working rather than an empty
  # prompt. An enumerate query is used deliberately: it is answered from SQL, so
  # the page is complete in milliseconds and the capture is deterministic
  # instead of racing a model.
  @{ name = "ask";        url = "$base/courses/$course/ask?q=Give+me+all+Chapter+3+questions+from+the+last+five+years" },
  @{ name = "analytics";  url = "$base/courses/$course/analytics" }
)

# preferredColorScheme: 1 = light, 2 = dark.
$scheme = if ($Theme -eq "dark") { 2 } else { 1 }

if ($Only.Count -gt 0) {
  $shots = $shots | Where-Object { $Only -contains $_.name }
}

foreach ($s in $shots) {
  $file = Join-Path $out "$($s.name).png"
  Write-Host "capturing $($s.name) ($Theme) ..."
  # Chrome reports the byte count on stderr even on success, so the strict
  # error preference has to be relaxed around the call itself.
  $prev = $ErrorActionPreference
  $ErrorActionPreference = "Continue"
  Start-Process -FilePath $chrome -Wait -NoNewWindow -ArgumentList @(
    "--headless=new", "--disable-gpu", "--hide-scrollbars",
    "--window-size=1440,1000", "--virtual-time-budget=8000",
    "--blink-settings=preferredColorScheme=$scheme",
    "--screenshot=$file", $s.url
  )
  $ErrorActionPreference = $prev
}

Get-ChildItem $out -Filter *.png | Select-Object Name, @{n='KB';e={[int]($_.Length/1KB)}}
