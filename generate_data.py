"""
Generate Fake Json Data
"""

import argparse
import json
import os
import random
from datetime import datetime, timedelta

def save_data(data:dict, current_date_str:str):
    filepath = os.path.normpath(os.path.join("./data",current_date_str + ".json"))
    with open(filepath, 'w', encoding = 'utf8') as f:
        json.dump(data, f, indent = 4)
        print(f"Finished saving:{current_date_str}")

# return object numbers / lat / lon
def generate_objects(max_position, max_num) -> dict:
    data = {"lons": [], "lats": [], "counts": []}
    # decide actual position number (min for 1)
    actual_position = random.randint(1, max_position)
    # generate fake data (lon & lat range inside Taiwan)
    for _ in range(actual_position):
        longitude = random.uniform(119, 112)
        latitude = random.uniform(20.5, 25.5)
        objects = random.randint(0, max_num)
        # Append to data
        data["lons"].append(longitude)
        data["lats"].append(latitude)
        data["counts"].append(objects)

    return data

def generate_dates(max_position, max_num, days, start_date) -> list:
    dates = []
    date_format = "%Y-%m-%d"
    current_date = datetime.strptime(start_date, date_format)

    for _ in range(days):
        current_date_str = current_date.strftime(date_format)
        current_date_data = generate_objects(max_position, max_num)
        save_data(data = current_date_data, current_date_str = current_date_str)
        dates.append(current_date_str)
        current_date += timedelta(days=1)

    return dates

def main():
    parser = argparse.ArgumentParser(description='Generate data based on input parameters.')
    parser.add_argument('--max_position', type=int, default = 30, help='Maximum number of positions to generate.')
    parser.add_argument('--max_num', type=int, default = 10, help='Maximum number of objects to generate.')
    parser.add_argument('--days', type=int, default = 30, help='Number of days between to generate.')
    parser.add_argument('--start_date', type=str, default='2022-01-01', help='Start date in the format YYYY-MM-DD.')

    args = parser.parse_args()

    dates = generate_dates(args.max_position,args.max_num, args.days, args.start_date)
    print("Generate:",dates)

if __name__ == "__main__":
    main()
