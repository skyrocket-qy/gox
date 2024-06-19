find . -type d -empty -exec touch {}/.gitkeep \; && \
echo "All empty directories now contain .gitkeep files."
