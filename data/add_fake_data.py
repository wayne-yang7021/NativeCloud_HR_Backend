from sqlalchemy import create_engine, text
from faker import Faker
import random
from datetime import datetime, timedelta
from dotenv import load_dotenv
import os

# === 連線設定 ===

# 載入 .env 檔案
load_dotenv()

# 組合 DB_URL
DB_HOST = os.getenv("DB_HOST")
DB_USER = os.getenv("DB_USER")
DB_PASSWORD = os.getenv("DB_PASSWORD")
DB_NAME = os.getenv("DB_NAME")
DB_PORT = os.getenv("DB_PORT")

DB_URL = f"postgresql://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}"
engine = create_engine(DB_URL)
fake = Faker()

# === 建立 Organization ===
def insert_organizations(conn, num_orgs=3):
    org_ids = []
    for _ in range(num_orgs):
        name = fake.company()
        result = conn.execute(text(
            "INSERT INTO organization (name) VALUES (:name) RETURNING organization_id"
        ), {"name": name})
        org_id = result.fetchone()[0]
        org_ids.append(org_id)
    print(f"✅ 已插入 {num_orgs} 個 organization")
    return org_ids

def insert_fixed_organizations(conn):
    organizations = [
        {"name": "CEO", "organization_id": "L1"},
        {"name": "COO", "organization_id": "L10"},
        {"name": "HR Manager", "organization_id": "L100"},
        {"name": "Operations Manager", "organization_id": "L101"},
        {"name": "CFO", "organization_id": "L11"},
        {"name": "Accounting Team", "organization_id": "L110"},
        {"name": "Finance Team", "organization_id": "L111"},
        {"name": "CTO", "organization_id": "L12"},
        {"name": "Engineering Team", "organization_id": "L120"},
        {"name": "IT Support", "organization_id": "L121"},
    ]

    for org in organizations:
        conn.execute(text("""
            INSERT INTO Organization (organization_id, name)
            VALUES (:organization_id, :name)
        """), org)
    
    print(f"✅ 已插入固定的 10 個 Organization！")


# === 建立 Employee ===
def insert_employees(conn, org_id, num_employees=5, create_manager=True):
    """
    為單一 org_id 插入員工資料（包含 Email 欄位）

    Args:
        conn: 資料庫連線
        org_id: 要插入的 organization_id
        num_employees: 插入的員工數量
        create_manager: 是否在第一位員工設定 is_manager=True
    """
    employee_ids = []
    manager_created = not create_manager  # 如果不需要 manager，預設已經有 manager

    for _ in range(num_employees):
        is_manager = False
        if not manager_created:
            is_manager = True
            manager_created = True

        first_name = fake.first_name()
        last_name = fake.last_name()
        email = fake.email()
        password = fake.password(length=12)

        result = conn.execute(text("""
            INSERT INTO Employee (first_name, last_name, is_manager, password, email, organization_id)
            VALUES (:first_name, :last_name, :is_manager, :password, :email, :organization_id)
            RETURNING employee_id
        """), {
            "first_name": first_name,
            "last_name": last_name,
            "is_manager": is_manager,
            "password": password,
            "email": email,
            "organization_id": org_id
        })
        employee_id = result.fetchone()[0]
        employee_ids.append(employee_id)
        if(len(employee_ids) == 10):
            print(f"✅ organization {org_id}, 員工 {employee_id}, is_manager={is_manager}, email={email}, password={password}")

    print(f"✅ organization {org_id} 已插入 {num_employees} 個 employee")
    return employee_ids



# === 建立 Access_log ===
def insert_access_logs(conn, employee_ids, num_days=30):
    """
    為每個員工插入打卡紀錄 (Access_log)

    Args:
        conn: 資料庫連線
        employee_ids: 要插入的 employee_id 列表
        num_days: 每個人要插入幾天的打卡紀錄
    """
    gate_names = ["AZ_door_1", "AZ_door_2", "AZ_door_3", "AZ_door_4", "AZ_door_5", "AZ_door_6", "AZ_door_7", "AZ_door_8", "AZ_door_9", "AZ_door_10"]

    for emp_id in employee_ids:
        for i in range(num_days):
            date = datetime.today().date() - timedelta(days=i)
            if date.weekday() >= 5:  # 5 是週六, 6 是週日
                continue
            # 上班打卡 (IN)
            check_in_time = datetime.combine(date, datetime.min.time()) + timedelta(hours=8, minutes=random.randint(0, 45))
            conn.execute(text("""
                INSERT INTO Access_log (employee_id, access_time, direction, gate_type, gate_name, access_result)
                VALUES (:employee_id, :access_time, 'IN', 'entry', :gate_name, 'Admitted')
            """), {
                "employee_id": emp_id,
                "access_time": check_in_time,
                "gate_name": random.choice(gate_names)
            })

            # 下班打卡 (OUT)
            check_out_time = datetime.combine(date, datetime.min.time()) + timedelta(hours=17, minutes=random.randint(20, 59))
            conn.execute(text("""
                INSERT INTO access_log (employee_id, access_time, direction, gate_type, gate_name, access_result)
                VALUES (:employee_id, :access_time, 'OUT', 'exit', :gate_name, 'success')
            """), {
                "employee_id": emp_id,
                "access_time": check_out_time,
                "gate_name": random.choice(gate_names)
            })

    print(f"✅ 每個 employee 各插入 {num_days} 天的 access_log（IN/OUT, entrance/exit）")


# === 建立 Report ===
# def insert_reports(conn, employee_ids, num_days=7):
#     for emp_id in employee_ids:
#         for i in range(num_days):
#             date = datetime.today().date() - timedelta(days=i)

#             check_in_time = datetime.combine(date, datetime.min.time()) + timedelta(hours=9, minutes=random.randint(0, 30))
#             check_out_time = check_in_time + timedelta(hours=8, minutes=random.randint(-30, 30))

#             total_hours = round((check_out_time - check_in_time).seconds / 3600, 2)
#             late_arrival = check_in_time.time() > datetime.strptime("09:00", "%H:%M").time()
#             early_departure = check_out_time.time() < datetime.strptime("18:00", "%H:%M").time()

#             conn.execute(text("""
#                 INSERT INTO report (
#                     employee_id, check_in_time, check_out_time,
#                     check_in_gate, check_out_gate, report_date,
#                     total_stay_hours, late_arrival_status, early_departure_status
#                 ) VALUES (
#                     :employee_id, :check_in_time, :check_out_time,
#                     :check_in_gate, :check_out_gate, :report_date,
#                     :total_stay_hours, :late_arrival_status, :early_departure_status
#                 )
#             """), {
#                 "employee_id": emp_id,
#                 "check_in_time": check_in_time,
#                 "check_out_time": check_out_time,
#                 "check_in_gate": "Main Entrance",
#                 "check_out_gate": "Main Entrance",
#                 "report_date": date,
#                 "total_stay_hours": total_hours,
#                 "late_arrival_status": late_arrival,
#                 "early_departure_status": early_departure
#             })
#     print(f"✅ 每個 employee 各插入 {num_days} 天的 report")

# === 主程式：自由調用 ===
if __name__ == "__main__":
    with engine.connect() as conn:
        # org_ids = insert_fixed_organizations(conn)
        # employee_id = insert_employees(conn, 'L1', num_employees=1, create_manager=True)
        # insert_access_logs(conn, employee_id, num_days=30)
        # employee_id = insert_employees(conn, 'L10', num_employees=1, create_manager=True)
        # insert_access_logs(conn, employee_id, num_days=30)
        # employee_id = insert_employees(conn, 'L100', num_employees=1, create_manager=True)
        # insert_access_logs(conn, employee_id, num_days=30)
        # employee_id = insert_employees(conn, 'L101', num_employees=1, create_manager=True)
        # insert_access_logs(conn, employee_id, num_days=30)
        # employee_id = insert_employees(conn, 'L11', num_employees=1, create_manager=True)
        # insert_access_logs(conn, employee_id, num_days=30)
        employee_id = insert_employees(conn, 'L110', num_employees=20, create_manager=False)
        insert_access_logs(conn, employee_id, num_days=50)
        employee_id = insert_employees(conn, 'L111', num_employees=20, create_manager=False)
        insert_access_logs(conn, employee_id, num_days=50)
        # employee_id = insert_employees(conn, 'L12', num_employees=1, create_manager=True)
        # insert_access_logs(conn, employee_id, num_days=30)
        employee_id = insert_employees(conn, 'L120', num_employees=20, create_manager=False)
        insert_access_logs(conn, employee_id, num_days=50)
        employee_id = insert_employees(conn, 'L121', num_employees=20, create_manager=False)
        insert_access_logs(conn, employee_id, num_days=50)
        # employee_id = ["781f6ff8-0ec6-44b9-9150-854160f02c82", 
        #                "ef3e1393-4977-470d-a72c-ccbebd55d158",
        #                "05f13451-51f7-4855-999d-1ee2ddcbc740",
        #                "5f2a2934-fdf3-4952-b695-428b8c277380",
        #                "4f39542c-d924-49ac-acde-652c3a156262",
        #                "aba4b238-8905-4f18-872b-83c16eb50a72",
        #                "f178a991-c659-4cff-82f3-33ad8fafb56d",
        #                "92a61758-6834-4cb1-8fa7-36a0596eee64",
        #                "81debf06-a18a-49e7-a8e8-52fcd3bda025",
        #                "b56b3d6a-e298-4e80-bc69-5032fdb9c29e",
        #                "efe6676b-04c0-400b-9db6-4b5df8d6b6b6",
        #                "9907e389-2e70-4864-bc1a-ca03edc74ffd",
        #                "fcaebeea-43dc-42d5-8a42-b11125cc8b01",
        #                "77a77d25-bd73-4d4d-bb65-5fef5d5b3ca2",
        #                "147889dc-711b-4a1e-96d8-f80f6590ca9e",
        #                "ed740955-1547-425f-a987-14b39baf974b",
        #                "4d08da02-a1ce-42d1-9540-5e4c6cdeca48",
        #                "0fb978e9-6db9-44db-9ecb-8dd5c541c848",
        #                "7281ce41-bb2b-4084-b5bd-507bb7a8d0ed",
        #                "a8c4c691-d360-4a34-a4ae-14a82b70102a",
        #                "4e8695ee-5a50-463a-9d2f-8216cadd309c",
        #                "39f0b626-c501-4433-858f-274bfe2dbeca",
        #                "dd62d5ba-018c-417b-ad67-21e50354e9a6",
        #                "d37856f0-b5b7-4bc9-bad4-0bbd328e2f84",
        #                "cf81680c-edd9-4ba4-a3e6-0e65996908ad",
        #                "791c7e38-8048-4695-ad9e-56fdd3ef9ad0",
        #                "b208426a-f8d9-414c-a6c0-b665172c258d",
        #                "3c28b0fd-7c5a-4b9e-a7bb-1e97cf5fea7d",
        #                "d9b2027a-ef0d-4c92-a5af-071552ea541f",
        #                "d8a6a5d2-aa48-40cf-9571-5fa36272c389",
        #                "25a24d12-f617-4741-977c-3a06f14ba186",
        #                "1f56db87-ef44-416d-abbe-71863489290c",
        #                "36abbb68-691a-4c11-bf6c-8acb290b85b6",
        #                "c116cdd8-1f82-493a-b7c7-6cc9520aa03d",
        #                "26f32a34-035c-45e3-aa7b-37d006a194a1",
        #                "33328b8f-81ac-4796-a415-6c811f2d075a",
        #                "c393a269-5ee7-484d-a2b3-b055a2a0ba3d",
        #                "44d6f45f-5433-45fc-8503-414339d39404",
        #                "cf52d1ab-12fd-4dc8-bdca-fc6a775a65bb",
        #                "170fa85d-4fe4-47fc-9757-e78d6b80585d",
        #                "314a6d64-c2a9-4736-a40b-bc68c579582e",
        #                "91e25336-4d9d-47e1-8bd5-146e9e620630",
        #                "47352caa-df6c-4f0e-89eb-7f927a45aa45",
        #                "96aa8124-3d1c-4233-bec3-c48fd1742946",
        #                "15d5b1a2-0ece-4cf2-8967-49e523027f82",
        #                "56ea2475-c67b-45cc-b728-b5a178f36101"
        #             ]
        # insert_access_logs(conn, employee_id, num_days=50)
        conn.commit()
    print("🎉 假資料全部插入完成！")