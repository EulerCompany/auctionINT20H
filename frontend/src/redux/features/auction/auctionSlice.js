
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from '../../../utils/axios'

const initialState = {
    user: null,
    token: null,
    isLoading: false,
    status: null,
}

export const createAuction = createAsyncThunk(
    'auction/createAuction', 
    async ({username, password}) => {
        try {
            const { data } = await axios.post('/auction/create', {
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

export const auctionSlice = createSlice({
    name: 'auction',
    initialState,
    reducers: {
        logout: (state) => {
            state.user = null
            state.token = null
            state.isLoading = false
            state.status = null
        }
    },
    extraReducers: builder => {
        builder.addCase(createAuction.pending, (state) => {
            state.isLoading = true
            state.status = null
        })
        builder.addCase(createAuction.fulfilled, (state, action) => {
            state.isLoading = false
            state.status = action.payload.message
            state.user = action.payload.user
            state.token = action.payload.token
        })
        builder.addCase(createAuction.rejected, (state, action) => {
            state.status = action.payload.message
            state.isLoading = false
        })
    },
})


// Why do we need this?
export default auctionSlice.reducer
