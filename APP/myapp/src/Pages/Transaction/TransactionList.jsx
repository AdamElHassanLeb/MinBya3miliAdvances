import React, {useContext, useEffect, useState} from 'react'
import {UserContext} from "../../utils/UserContext";
import TransactionService from "../../services/TransactionService";
import {Box, FormControl, FormControlLabel, FormLabel, Radio, RadioGroup, Typography} from "@mui/material";
import TransactionListItem from "../../components/Transaction/TransactionListItem";

const TransactionList = () => {

    const [status, setStatus] = useState("All"); // "accepted", "rejected", "pending"
    const [direction, setDirection] = useState("incoming"); // "incoming", "outgoing"

    const {user} = useContext(UserContext);
    const [transactions, setTransactions] = useState([]);


    useEffect(() => {

       handleChange()
    }, [status, direction]);


    const handleStatusChange = (event) => {
        setStatus(event.target.value);
        handleChange();
    };

    const handleDirectionChange = (event) => {
        setDirection(event.target.value);
        handleChange();
    };


    const handleChange = async () => {
        try{
            if(direction === "incoming"){
                const transactions = await TransactionService.getTransactionsByOfferedUserAndStatus(user.user_id, status)
                setTransactions(transactions);
            }else{
                const transactions = await TransactionService.getTransactionsByOfferingUserAndStatus(user.user_id, status)
                setTransactions(transactions);
            }
        }catch(err){
            console.error(err);
            setTransactions([])
        }

    }

    return (<>
        <Box
            className="modal-people"
            sx={{
                position: 'absolute',
                top: '5%',
                left: '50%',
                transform: 'translate(-50%, 0)',
                width: '80%',
                height: '70%',
                borderRadius: 2,
                boxShadow: 24,
                p: 4,
                display: 'flex',
                flexDirection: 'column',
                marginTop: '10vh'
            }}
        >


            <Box>
                <FormControl>
                    <FormLabel>Direction</FormLabel>
                    <RadioGroup value={direction} onChange={handleDirectionChange} row>
                        <FormControlLabel value="incoming" control={<Radio/>} label="Incoming"/>
                        <FormControlLabel value="outgoing" control={<Radio/>} label="Outgoing"/>
                    </RadioGroup>
                </FormControl>

                <FormControl>
                    <FormLabel>Status</FormLabel>
                    <RadioGroup value={status} onChange={handleStatusChange} row>
                        <FormControlLabel value="Accepted" control={<Radio/>} label="Accepted"/>
                        <FormControlLabel value="Completed" control={<Radio/>} label="Completed"/>
                        <FormControlLabel value="Pending" control={<Radio/>} label="Pending"/>
                        <FormControlLabel value="All" control={<Radio/>} label="All"/>
                    </RadioGroup>
                </FormControl>
            </Box>


            {
                transactions && transactions.length > 0 ? (<>
                        <Typography>Transactions:</Typography>
                        <ul style={{
                            padding: 0,
                            listStyle: "none",
                            margin: 0,
                            background: "transparent",
                            marginTop: 20,
                            overflowY: "auto",
                        }}>
                            {
                                transactions.map((item, index) => (
                                    <li key={index} style={{marginBottom: "16px", background: "transparent"}}>
                                        <TransactionListItem transaction={item}/>
                                    </li>
                                ))
                            }
                        </ul>
                    </>
                ) : (<Typography> No Transactions </Typography>)
            }

        </Box>

    </>)
}

export default TransactionList