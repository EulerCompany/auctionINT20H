
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from '../../../utils/axios'

export const createAuction = createAsyncThunk(
    'auctions/createAuction', 
    async ({title, description, start_price, start_date, end_date, files}) => {

        try {
            const { data } = await axios.post('/auctions/create', {
                title, 
                description, 
                start_price, 
                start_date, 
                end_date,
                files
            })
            // TODO: remove log
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
})

export const updateAuction = createAsyncThunk(
    'auctions/createAuction', 
    async ({title, description, id}) => {

        try {
            const { data } = await axios.patch(`/auctions/${id}/update`, {
                title, 
                description
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
            const { data } = await axios.get('/auctions/active')
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
            const { data } = await axios.get(`/users/${id}/auction/active`)
            // TODO: remove log
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
    }
)

export const fetchBetsForAuction = createAsyncThunk(
    'auctions/fetchBetsForAuction',
    async (id) => {
        try {
            const { data } = await axios.get(`/auctions/${id}/bets`)
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
    }
)

export const makeBet = createAsyncThunk(
    'auctions/makeBet', 
    async ({id, bet}) => {
        console.log("AHTUNG" ,id, bet)
        try {
            const { data } = await axios.post(`/auctions/${id}/makebet`, {
                bet
            })
            // TODO: remove log
            console.log(data)
            return data
            
        } catch (error) {
            console.log(error)
        }
})



export const fetchAuctionPhotos = createAsyncThunk(
    'auctions/fetchAuctionPhotos',
    async (id) => {
        try {
            const { data } = await axios.get(`/auctions/${id}/photos`)
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
        total: 0,
        auctions: [],
        bets: [{},],
        photos: [],
    },
    reducers: {
    },
    extraReducers: builder => {
        // all auctions
        builder.addCase(fetchAllAuctions.pending, (state) => { 
            console.log("pending.... at fetchall")
            state.loading = true
        })
        builder.addCase(fetchAllAuctions.fulfilled, (state, action) => {
            console.log("fulfilled")
            state.loading = false
            state.auctions = action.payload
            state.total = action.payload.length
            
        })
        builder.addCase(fetchAllAuctions.rejected, (state, action) => {
            console.log("rejected")
            state.loading = false
        })
        // auctions by user
        builder.addCase(fetchAllAuctionsByUserId.pending, (state) => { 
            console.log("pending.... at fetchall")
            state.loading = true
        })
        builder.addCase(fetchAllAuctionsByUserId.fulfilled, (state, action) => {
            console.log("fulfilled")
            state.loading = false
            state.auctions = action.payload
            state.total = action.payload.length
            
        })
        builder.addCase(fetchAllAuctionsByUserId.rejected, (state, action) => {
            console.log("rejected")
            state.loading = false
        })



        builder.addCase(fetchBetsForAuction.pending, (state) => { 
            console.log("pending.... at fetchall")
            state.loading = true
        })
        builder.addCase(fetchBetsForAuction.fulfilled, (state, action) => {
            console.log("fulfilled")
            state.bets = action.payload
            state.loading = false
        })
        builder.addCase(fetchBetsForAuction.rejected, (state, action) => {
            console.log("rejected")
            state.loading = false
        })


        // photos
        builder.addCase(fetchAuctionPhotos.pending, (state) => { 
            console.log("pending auction photos")
            // state.loading = true
        })
        builder.addCase(fetchAuctionPhotos.fulfilled, (state, action) => {
            console.log("fulfilled auction photos")
            // state.loading = false
            const base64Images = action.payload.map(item => `data:image/png;base64,${item.img_base64}`);
            state.photos = base64Images;
        })
        builder.addCase(fetchAuctionPhotos.rejected, (state, action) => {
            console.log("rejected auction photos")
            // state.loading = false
        })
}})


// Why do we need this?
export default auctionSlice.reducer
