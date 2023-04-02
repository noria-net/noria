#!/bin/bash

# Make sure GITHUB_API_TOKEN is set
if [ -z "$GITHUB_API_TOKEN" ]; then
	echo "Error: please provide a GITHUB_API_TOKEN for api access"
	exit 1
fi

# Get the release version
TAG=$(git describe --tags --abbrev=0)

# Create a new release
curl -L \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_API_TOKEN"\
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/noria-net/noria/releases \
  -d @- << EOF
{
  "tag_name": "$TAG",
  "target_commitish": "main",
  "name": "$TAG",
  "draft": true,
  "prerelease": false,
  "generate_release_notes": true
}
EOF

# Make sure that release was created
if [[ $? -ne 0 ]]; then
	echo "Call to create release failed"
	exit 1
fi


# Define variables
FILENAME=noriad
VERSION=$(./noriad --version)
SHA256=$(shasum -a 256 $FILENAME | cut -d ' ' -f1)
RELEASE_NOTES="Release $VERSION"$'\n\n'"SHA256: $SHA256"
USERNAME=jumpdest7d
REPO=noria-net

# Zip up binary
zip $FILENAME.zip $FILENAME

# Create release on GitHub
curl -H "Authorization: token YOUR_GITHUB_TOKEN" \
     -H "Content-Type:application/json" \
     -X POST \
     -d "{\"tag_name\": \"$VERSION\", \"name\": \"$VERSION\", \"body\": \"$RELEASE_NOTES\"}" \
     https://api.github.com/repos//YOUR_REPO/releases

# Upload zip file to release
curl -H "Authorization: token YOUR_GITHUB_TOKEN" \
     -H "Content-Type: application/zip" \
     --data-binary @$FILENAME.zip \
     "https://uploads.github.com/repos/YOUR_USERNAME/YOUR_REPO/releases/$VERSION/assets?name=$FILENAME.zip"

