import os
import time
from datetime import datetime, timedelta

def list_files_modified_days_ago(directory, days_ago=360, recursive=True):
    target_date = datetime.now() - timedelta(days=days_ago)
    target_day_start = datetime.combine(target_date.date(), datetime.min.time())
    target_day_end = datetime.combine(target_date.date(), datetime.max.time())
    
    print(f"Looking for files modified on: {target_date.strftime('%Y-%m-%d')}")

    matched_files = []

    for root, dirs, files in os.walk(directory):
        for file in files:
            filepath = os.path.join(root, file)
            mtime = datetime.fromtimestamp(os.path.getmtime(filepath))
            if target_day_start <= mtime <= target_day_end:
                matched_files.append(filepath)

        if not recursive:
            break

    return matched_files

# Example usage
if __name__ == "__main__":
    directory_to_search = "."  # current directory
    files = list_files_modified_days_ago(directory_to_search, days_ago=7)
    
    if files:
        print("\nFiles modified 7 days ago:")
        for file in files:
            print(file)
    else:
        print("No files found modified exactly 7 days ago.")