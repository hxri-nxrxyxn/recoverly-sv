<script>
    import { Link, useRouter } from "svelte-routing";
    import Morning from "./assets/undraw_mornings_kmib.png";
    import { Storage } from "@capacitor/storage";
    import Trophy from "./assets/trophy-star-2.svg";
    import Nav from "./Nav.svelte";

    let first = $state(0);
    let second = $state(0);
    const decFirst = () => {
        if (first != 0) first--;
    };
    const incFirst = () => {
        first++;
    };
    const decSecond = () => {
        if (second != 0) second--;
    };
    const incSecond = () => {
        second++;
    };

    // navigation

    const router = useRouter();
    let hasToken = true; // Start as null to represent "loading"

    let claims;

    const logout = async () => {
        try {
            await Storage.remove({ key: "id" });
            location.href = "/login";
        } catch (error) {
            console.error("Error:", error);
        }
    };

    const checkUser = async () => {
        try {
            const value = await Storage.get({ key: "id" });
            const id = value.value;
            if (!id) {
                hasToken = false;
                location.href = "/login";
            }
            const response = await fetch(
                `http://192.168.183.224:8080/api/v1/users/${id}`,
                {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                    },
                },
            );

            if (response.ok) {
                const data = await response.json();
                console.log(data);
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };

    checkUser();
</script>

{#if hasToken === true}
    <main>
        <Nav />
        <div class="pills">
            <div class="pill">
                <p>NEW</p>
            </div>
            <Link to="/events">
                <div class="pill">
                    <p>EVENTS</p>
                </div>
            </Link>
            <Link to="/chatroom">
                <div class="pill">
                    <p>CHATROOM</p>
                </div>
            </Link>
        </div>
        <h1>Good morning, <span>Hari</span></h1>

        <img src={Morning} alt="" class="hero" />

        <div class="card">
            <h1>Daily Task</h1>
            <div class="tasks">
                <div class="task">
                    <div class="task__text">
                        <h5>No. of minutes meditated</h5>
                        <p>0/15</p>
                    </div>
                    <div class="task__adjust">
                        <button onclick={decFirst}>-</button>
                        <h5>{first}</h5>
                        <button onclick={incFirst}>+</button>
                    </div>
                </div>
                <div class="task">
                    <div class="task__text">
                        <h5>No. of cigarretes smoked</h5>
                        <p>0/7</p>
                    </div>
                    <div class="task__adjust">
                        <button onclick={decSecond}>-</button>
                        <h5>{second}</h5>
                        <button onclick={incSecond}>+</button>
                    </div>
                </div>
                <button
                    onclick={alert(
                        `that's ${first < 15 && second < 7 ? "very nice of you :)" : "very bad for your health!"}`,
                    )}>submit</button
                >
            </div>
        </div>

        <div class="achievements">
            <h2>Achivements</h2>
            <div class="achievement">
                <div class="achievement__logo">
                    <img src={Trophy} alt="" class="icon" />
                </div>
                <div class="achievement__text">
                    <p>Fresh Start</p>
                </div>
            </div>
            <div class="achievement">
                <div class="achievement__logo">
                    <img src={Trophy} alt="" class="icon" />
                </div>
                <div class="achievement__text">
                    <p>Smoke Free 2nd Day</p>
                </div>
            </div>
        </div>

        <div class="emergency">
            <Link to="/emergency">SOS</Link>
        </div>
    </main>
{:else if hasToken === false}
    <p>Redirecting to login...</p>
{:else}
    <p>Loading...</p>
{/if}

<style>
    /* Your styles */
    .task {
        display: flex;
        justify-content: space-between;
        margin-top: 1rem;
    }
    .task__adjust {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .task__adjust > * {
        margin-left: 1rem;
    }
    .task__adjust button {
        position: initial;
        font-size: 1rem;
        padding: 0.5rem;
        border: 0;
        color: var(--color-text);
        background-color: var(--color-light);
    }
    .task__adjust h5 {
        text-align: center;
    }
    .task__text p {
        font-size: 0.75rem;
    }
    .pills {
        display: flex;
        margin-bottom: 1rem;
    }
    .pill {
        background-color: var(--color-light);
        padding: 0.25rem 1rem;
        margin-right: 0.5rem;
        border-radius: 2rem;
    }
    .pill p {
        font-size: 0.5rem;
        font-weight: 700;
        color: var(--color-primary);
    }
    .achievements {
        margin-top: 2rem;
    }
    .achievement {
        display: flex;
        margin: 1rem 0;
    }
    .achievement__text {
        margin-left: 1rem;
    }
    .achievement__text p {
        font-weight: 600;
    }
    .achievement__logo img {
        transform: translateY(0.2rem);
    }
</style>
