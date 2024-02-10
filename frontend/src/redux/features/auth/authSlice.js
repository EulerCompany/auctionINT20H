import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from '../../../utils/axios'


export const registerUser = createAsyncThunk(
    'auth/registerUser', 
    async ({username, password}) => {
        try {
            const { data } = await axios.post('/user/signup', {
                username, 
                password
            })
            // TODO: remove log
            console.log(data)
            if(data.token) {
                window.localStorage.setItem('token', data.token)
            }

            return data
            
        } catch (error) {
            console.log(error)
        }
})

export const loginUser = createAsyncThunk(
    'auth/loginUser', 
    async ({username, password}) => {
        try {
            
            const { data } = await axios.post('/user/login', {
                username,
                password,
            })

            if(data.token) {
                window.localStorage.setItem('token', data.token)
            }

            return data
            
        } catch (error) {
            console.log(error)
        }
})

export const getMe = createAsyncThunk(
    'auth/getMe', 
    async () => {
        try {
            
            const { data } = await axios.get('/auth')
            return data
            
        } catch (error) {
            console.log(error)
        }
})


export const authSlice = createSlice({
    name: 'auth',
    initialState:{
    user: null,
    token: null,
    isloading: false,
    status: null,
    userId: null,
    },
    reducers: {
        logout: (state) => {
            state.user = null
            state.token = null
            state.isLoading = false
            state.status = null
        }
    },
    extraReducers: builder => {
        builder.addCase(registerUser.pending, (state) => {
            state.isLoading = true
            state.status = null
        })
        builder.addCase(registerUser.fulfilled, (state, action) => {
            state.isLoading = false
            state.status = action.payload.message
            state.user = action.payload.user
            state.token = action.payload.token
            state.userId = action.payload.userId
        })
        builder.addCase(registerUser.rejected, (state, action) => {
            state.status = action.payload.message
            state.isLoading = false
        })

        builder.addCase(loginUser.pending, (state) => {
            state.isLoading = true
            state.status = null
        })
        builder.addCase(loginUser.fulfilled, (state, action) => {
            state.isLoading = false
            state.status = action.payload.message
            state.user = action.payload.user
            state.token = action.payload.token
            state.userId = action.payload.userId
        })
        builder.addCase(loginUser.rejected, (state, action) => {
            state.status = action.payload.message
            state.isLoading = false
        })


        builder.addCase(getMe.pending, (state) => {
            state.isLoading = true
            state.status = null
        })
        builder.addCase(getMe.fulfilled, (state, action) => {
            state.isLoading = false
            state.status = null
            state.user = action.payload?.user
            state.token = action.payload?.token
        })
        builder.addCase(getMe.rejected, (state, action) => {
            state.status = action.payload.message
            state.isLoading = false
        })
    },
})

export const isAuth = state => Boolean(state.auth.token)

export const { logout } = authSlice.actions

export default authSlice.reducer
