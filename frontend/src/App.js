import React, { useState } from 'react';
import { Button, TextField, List, ListItem, ListItemText, Typography, Box } from '@mui/material';

function App() {
    const [packs, setPacks] = useState({});
    const [itemsOrdered, setItemsOrdered] = useState(8750); // Initial value is set to 10
    const [error, setError] = useState(null);

    const fetchData = (e) => {
        e.preventDefault(); // Prevent the page from refreshing

        if (itemsOrdered <= 0) {
            setPacks({});
            setError('Value should be greater than zero.');
            return;
        }

        fetch(`http://localhost:8080/api/v1/get-packs-number/${itemsOrdered}`)
            .then(response => {
                if (!response.ok) {
                    return response.text().then(text => {
                        setPacks({});
                        throw new Error(`Error: ${response.status} - ${response.statusText} - ${text}`);
                    });
                }
                return response.json();
            })
            .then(data => {
                setPacks(data);
                setError(null);
            })
            .catch(error => {
                console.error('Error fetching data:', error);
                setError(error.message);
            });
    };

    return (
        <Box sx={{ maxWidth: 400, m: 'auto', p: 2 }}>
            <h1>Packs calculator</h1>

            <Box>
                <form onSubmit={fetchData}>
                    <TextField
                        type="number"
                        label="Items Ordered"
                        value={itemsOrdered}
                        onChange={(e) => setItemsOrdered(e.target.value)}
                        variant="outlined"
                        fullWidth
                        required
                        inputProps={{ required: true }}
                        sx={{ mb: 1 }}
                    />
                    <Button variant="contained" type="submit">
                        Fetch Packs
                    </Button>
                </form>
            </Box>
            <Box sx={{ mt: 2 }}>
                <Typography variant="h6">Result:</Typography>
                {error && (
                    <Typography variant="body1" color="error" sx={{ mt: 2 }}>
                        Error: {error}
                    </Typography>
                )}
                <List>
                    {Object.keys(packs).map((key, index) => (
                        <ListItem key={index}>
                            <ListItemText primary={`${key}: ${packs[key]}`} />
                        </ListItem>
                    ))}
                </List>
            </Box>
        </Box>
    );
}

export default App;
