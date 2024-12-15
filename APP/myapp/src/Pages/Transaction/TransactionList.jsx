import React from 'react'
import {useContext, useEffect, useState} from "react";
import {UserContext} from "../../utils/UserContext";
import TransactionService from "../../services/TransactionService";
import {Box, Typography} from "@mui/material";
import TransactionListItem from "../../components/Transaction/TransactionListItem";


const TransactionList = () => {

    const { user } = useContext(UserContext);
    const [transactions, setTransactions] = useState([]);


    useEffect(() => {

        async function fetchTransactions() {
            try {
                const transactionsData = await TransactionService.getTransactionsByOfferedUserAndStatus(user.user_id, "");
                setTransactions(transactionsData)
            } catch (e) {
                console.error(e);

            }
        }

        fetchTransactions();
    }, []);


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

            {
                transactions && transactions.length > 0 ? (

                    <ul style={{ padding: 0, listStyle: "none", margin: 0, background: "transparent" }}>
                        {
                            transactions.map((item, index) => (
                                <li key={index} style={{ marginBottom: "16px", background: "transparent" }}>
                                    <TransactionListItem transaction={item} />
                                </li>
                            ))
                        }
                    </ul>

                ) : (<Typography> No Transactions </Typography>)
            }

        </Box>

    </>)
}

export default TransactionList