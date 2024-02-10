
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from '../../../utils/axios'

export const createAuction = createAsyncThunk(
    'auctions/createAuction', 
    async ({username, password}) => {
        try {
            const { data } = await axios.post('/auctions/create', {
                username, 
                password
            })
            // TODO: remove log
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
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
)

export const fetchAllAuctionsByUserId = createAsyncThunk(
    'auctions/fetchAllAuctionsByUserId',
    async (id) => {
        try {
            const { data } = await axios.get(`/user/${id}/auction/active`)
            // TODO: remove log
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
    }
)

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
            state.totalPages = 123
            state.loading = false
        })
        builder.addCase(fetchAllAuctions.rejected, (state, action) => {
            console.log("rejected")
            state.loading = false
        })
        builder.addCase(fetchAllAuctionsByUserId.pending, (state) => { 
            console.log("pending.... at fetchall")
            state.loading = true
        })
        builder.addCase(fetchAllAuctionsByUserId.fulfilled, (state, action) => {
            console.log("fulfilled")
            state.auctions = action.payload
            state.totalPages = 123
            state.loading = false
        })
        builder.addCase(fetchAllAuctionsByUserId.rejected, (state, action) => {
            console.log("rejected")
            state.loading = false
        })
}})


// Why do we need this?
export default auctionSlice.reducer
