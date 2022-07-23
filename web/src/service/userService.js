import request from '../utils/request';

const register = ({ username, telephone, password }) => request.post('auth/register', { username, telephone, password });
const login = ({ telephone, password }) => request.post('auth/login', { telephone, password });
const info = () => request.get('auth/info');
export default {
  register,
  login,
  info,
};
