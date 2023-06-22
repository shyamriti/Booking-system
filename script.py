import csv
import sqlite3

def create_connection():
    conn = None
    try:
        conn = sqlite3.connect('new.db')  
        return conn
    except sqlite3.Error as e:
        print(e)
    
    return conn


def create_table(conn):
    try:
        cursor = conn.cursor()
        cursor.execute('''CREATE TABLE IF NOT EXISTS seat_pricings
                          (id INTEGER, seat_class TEXT, min_price TEXT, normal_price TEXT, max_price TEXT)''')  
        cursor.execute('''CREATE TABLE IF NOT EXISTS seats
                          (seat_id INTEGER, seat_identifier TEXT, seat_class TEXT)''')  
    except sqlite3.Error as e:
        print(e)

def insert_idata_in_seatpricings(conn, table_name, data):
    try:
        cursor = conn.cursor()
        cursor.execute(f"INSERT INTO {table_name} VALUES (?, ?, ?, ?, ?)", data)  
        conn.commit()
    except sqlite3.Error as e:
        print(e)

def insert_data_in_seats(conn, table_name, data):
    try:
        cursor = conn.cursor()
        cursor.execute(f"INSERT INTO {table_name} VALUES (?, ?, ?)", data)  
        conn.commit()
    except sqlite3.Error as e:
        print(e)

def main():
    conn = create_connection()
    if conn is not None:
        create_table(conn)

        with open('/home/shyam/shyamworkspace/Booking-service/pkg/files/seat_pricings.csv', 'r') as file1:  
            csv_data1 = csv.reader(file1)
            next(csv_data1)  
            for row in csv_data1:
                insert_idata_in_seatpricings(conn, 'seat_pricings', row)

        with open('/home/shyam/shyamworkspace/Booking-service/pkg/files/seats.csv', 'r') as file2:  
            csv_data2 = csv.reader(file2)
            next(csv_data2)  
            for row in csv_data2:
                insert_data_in_seats(conn, 'seats', row)

        conn.close()
        print("Data uploaded successfully!")
    else:
        print("Error connecting to the database.")

if __name__ == '__main__':
    main()
