export const api_server_url = "http://localhost:8080/api";

export const auth_server_url = api_server_url + "/auth";

interface RequestInitType {
  mode: 'cors' | 'no-cors';
}

export const request_init: RequestInitType = {
  mode: 'cors',
};