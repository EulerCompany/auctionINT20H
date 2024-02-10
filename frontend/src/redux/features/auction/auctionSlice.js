
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from '../../../utils/axios'

export const createAuction = createAsyncThunk(
    'auctions/createAuction', 
    async ({username, password}) => {
        try {
            const { data } = await axios.post('/auction/create', {
                username, 
                password
            })
            // TODO: remove log
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
}, {
    pending: (state) => {},
    rejected: (state, action) => {},
    fulfilled: (state, action) => {},

})

export const fetchAllAuctions = createAsyncThunk(
    'auctions/fetchAllAuctions',
    async () => {
        try {
            const { data } = await axios.get('/auction/active')
            // TODO: remove log
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
    }
, { // why this isn't working?
    pending: (state) => {
    },
    rejected: (state, action) => {
    },
    fulfilled: (state, action) => {
    },
})

export const auctionSlice = createSlice({
    name: 'auctions',
    initialState: {
        loading: false,
        pageSize: 10,
        totalPages: 0,
        auctions: []
    },
    reducers: {
    },
    extraReducers: builder => {
        builder.addCase(fetchAllAuctions.pending, (state) => { 
            console.log("pending.... at fetchall")
            state.loading = true
        })
        builder.addCase(fetchAllAuctions.fulfilled, (state, action) => {
            console.log("fulfilled")
            state.auctions = action.payload
            state.totalPages = action.payload
            state.loading = false
        })
        builder.addCase(fetchAllAuctions.rejected, (state, action) => {
            console.log("rejected")
            state.loading = false
        })
}})


// Why do we need this?
export default auctionSlice.reducer
