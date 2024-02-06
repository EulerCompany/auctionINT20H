import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://localhost:8080'
})

instance.interceptors.request.use(config => {
    // TODO: do we actually need to prepend bearer with JWT?
    config.headers.Authorization = "Bearer " + window.localStorage.getItem('token')
    return  config
})

export default instance
