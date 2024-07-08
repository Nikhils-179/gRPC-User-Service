from python_pb_files import user_pb2 as user_pb
from python_pb_files import user_pb2_grpc as user_grpc
from google.protobuf import empty_pb2 
import logging
import grpc


logging.basicConfig(format='%(asctime)s:%(name)s : %(message)s', level=logging.DEBUG)
log = logging.getLogger(__name__)

socket = 'localhost:50051'
channel = grpc.insecure_channel(socket)
client = user_grpc.UserServiceStub(channel)
log.info(f"Connected to the Server : {socket}")

def call_get_user(id: str):
    log.info(f"Calling GetUser with the user id {id}")
    user_id_message = user_pb.UserID()    
    user_id_message.id = id
    try:
        res = client.GetUser(user_id_message)
        log.info("Response")
        log.info(res)
    except grpc.RpcError as e:
        log.error(f"Errored while getting the user with ID {id}")
        log.error(e)

def call_create_user(id: str, first_name: str, last_name: str, city: str, phone: str, height: float, married: bool):
    log.info(f"Creating user with the user id {id}")

    req_pb = user_pb.User(
        id=id,
        first_name=first_name,
        last_name=last_name,
        city=city,
        phone=phone,
        height=height,
        married=married
    )
    try:
        res = client.SetUser(req_pb)
        log.info(f"Successfully created the User with the ID {res.id}")
    except grpc.RpcError as e:
        log.error(f"Errored while creating the user with ID {id}")
        log.error(e)

def call_list_users():
    log.info("Calling ListUsers")
    try:
        res = client.ListUsers(empty_pb2.Empty())
        log.info("ListUsers Results")
        for r in res.users:
            log.info(r)
    except grpc.RpcError as e:
        log.error("Errored while listing users")
        log.error(e)

def call_update_user(id: str, first_name: str, last_name: str, city: str, phone: str, height: float, married: bool):
    log.info(f"Calling UpdateUser with the user id {id}")
    req_pb = user_pb.User(
        id=id,
        first_name=first_name,
        last_name=last_name,
        city=city,
        phone=phone,
        height=height,
        married=married
    )
    try:
        client.UpdateUser(req_pb)
        log.info("Successfully updated")
    except grpc.RpcError as e:
        log.error(f"Errored while updating the user with ID {id}")
        log.error(e)

def call_delete_user(id: str):
    log.info(f"Calling DeleteUser with the user id {id}")
    try:
        client.DeleteUser(user_pb.UserID(id=id))
        log.info(f"Successfully deleted the user with ID {id}")
    except grpc.RpcError as e:
        log.error(f"Errored while deleting the user with ID {id}")
        log.error(e)

def call_search_users(city: str, phone: str, married: bool):
    log.info(f"Calling SearchUsers with the criteria city={city}, phone={phone}, married={married}")
    criteria = user_grpc.SearchCriteria(City=city, Phone=phone, Married=married)
    try:
        res = client.SearchUsers(criteria)
        log.info("SearchUsers Results")
        for r in res.Users:
            log.info(r)
    except grpc.RpcError as e:
        log.error("Errored while searching users")
        log.error(e)

def call_get_users_by_id(user_ids):
    log.info(f"Calling GetUsersByID with the IDs: {user_ids}")
    ids = user_pb.UserIDs(Ids=user_ids)
    try:
        res = client.GetUsersByID(ids)
        log.info("GetUsersByID Results")
        for r in res.Users:
            log.info(r)
    except grpc.RpcError as e:
        log.error("Errored while getting users by IDs")
        log.error(e)

# Test calls
call_get_user("1")
call_delete_user(id="1")
call_delete_user(id="10")
call_create_user(id="2", first_name="Ram", last_name="Doe", city="SF", phone="5555555555", height=5.5, married=False)
call_get_user("10")
call_create_user(id="11", first_name="Bob", last_name="Khan", city="NY", phone="1234567890", height=6.1, married=True)
call_update_user(id="2", first_name="Bob", last_name="Smith", city="LA", phone="1234567890", height=6.1, married=True)
call_list_users()
call_delete_user(id="2")
call_update_user(id="2", first_name="Bob", last_name="Smith", city="LA", phone="1234567890", height=6.1, married=True)  # Should error

