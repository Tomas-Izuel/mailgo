package user

import (
	"context"
)

func GetUserData(userId string, ctx context.Context) (*UserResponse, error) {
	//baseURL := fmt.Sprintf("http://localhost:3000/v1/users/%s", userId)

	//req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	//if err != nil {
	//	return nil, fmt.Errorf("error creating HTTP request: %w", err)
	//}

	//req.Header.Set("Content-Type", "application/json")

	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	return nil, fmt.Errorf("error making HTTP request: %w", err)
	//}
	//defer resp.Body.Close()

	//if resp.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	//}

	//var userResponse UserResponse
	//if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
	//	return nil, fmt.Errorf("error decoding response JSON: %w", err)
	//}

	mockedEmail := UserResponse{
		ID:    userId,
		Name:  "Tomas Izuel",
		Email: "tomasizuelbackup@gmail.com",
	}

	return &mockedEmail, nil
}
