import pandas as pd
import numpy as np
import random
from concurrent.futures import ProcessPoolExecutor
import multiprocessing
import os

# ============================
# GLOBAL (dipakai multiprocessing)
# ============================
# Variabel global y, n_samples, n_actuators akan diinisialisasi
# di setiap proses anak menggunakan 'initializer'
y_global = None
n_samples_global = 0
n_actuators_global = 0
MUTATION_RATE = 0.1

# Fungsi inisialisasi untuk ProcessPoolExecutor
def initializer_worker(y_data, n_s, n_a):
    global y_global
    global n_samples_global
    global n_actuators_global
    y_global = y_data
    n_samples_global = n_s
    n_actuators_global = n_a

# ============================
# GA FUNCTIONS (HARUS GLOBAL)
# ============================
def create_chromosome():
    # Menggunakan n_actuators_global
    return np.random.randint(0, 2, n_actuators_global)

def create_population(pop_size):
    return [create_chromosome() for _ in range(pop_size)]

def evaluate_fitness(chromosome):
    # Menggunakan variabel global yang diinisialisasi per proses
    pred = np.tile(chromosome, (n_samples_global, 1))
    
    # Menghitung kecocokan
    matches = (pred == y_global).sum()
    
    # Pembagi yang aman: Pastikan tidak membagi dengan nol
    denominator = n_samples_global * n_actuators_global
    if denominator == 0:
        return 0.0 # Hindari pembagian dengan nol
    
    return matches / denominator

def selection(pop, fitness):
    i, j = random.sample(range(len(pop)), 2)
    return pop[i] if fitness[i] > fitness[j] else pop[j]

def crossover(p1, p2):
    # Menggunakan n_actuators_global
    point = random.randint(1, n_actuators_global - 1)
    return np.concatenate([p1[:point], p2[point:]])

def mutate(chrom):
    # Menggunakan n_actuators_global
    for i in range(n_actuators_global):
        if random.random() < MUTATION_RATE:
            chrom[i] ^= 1
    return chrom


# ============================
# MAIN
# ============================
if __name__ == "__main__":

    print("Current dir:", os.getcwd())
    # ... (kode loading data tetap sama) ...
    
    # ============================
    # LOAD DATA
    # ============================
    # PASTIKAN FILE "Dataset_screw/Filling_ALL.module.csv" ADA
    # di direktori kerja atau path-nya benar
    # Jika tidak, ini akan gagal sebelum multiprocessing dimulai.
    try:
        df = pd.read_csv("Dataset_screw/Filling_ALL.module.csv")
    except FileNotFoundError:
        print("ERROR: File 'Dataset_screw/Filling_ALL.module.csv' not found. Check your file path.")
        exit()
        
    print("Loaded:", df.shape)

    # ============================
    # SENSOR & ACTUATOR
    # ============================
    sensor_cols = [
        c for c in df.columns
        if c.startswith("I_") and set(df[c].dropna().unique()).issubset({0, 1})
    ]
    action_cols = [c for c in df.columns if c.startswith("O_")]

    df[sensor_cols] = df[sensor_cols].fillna(0).astype(int)
    df[action_cols] = df[action_cols].fillna(0).astype(int)

    # ============================
    # DATA MATRIX
    # ============================
    X = df[sensor_cols].values
    y = df[action_cols].values # Ini adalah target global (diganti namanya jadi y_main)
    
    # Ganti variabel global yang lama dengan yang baru
    n_samples_main = X.shape[0]
    n_actuators_main = y.shape[1]

    print("Jumlah data     :", n_samples_main)
    print("Jumlah sensor   :", X.shape[1])
    print("Jumlah actuator :", n_actuators_main)

    # ============================
    # GA CONFIG
    # ============================
    POP_SIZE = 40
    GENERATIONS = 60
    N_CORES = min(4, multiprocessing.cpu_count())

    print("GA parallel pakai core:", N_CORES)
    
    # SET VARIABEL GLOBAL DI MAIN UNTUK FUNGSI NON-PARALEL
    # Note: Meskipun sudah pakai initializer, fungsi create_chromosome/population
    # yang dipanggil sebelum loop GA masih butuh inisialisasi awal.
    n_actuators_global = n_actuators_main

    # ============================
    # GA LOOP
    # ============================
    population = create_population(POP_SIZE)

    for gen in range(GENERATIONS):
        # MENGGUNAKAN INITIALIZER
        with ProcessPoolExecutor(
            max_workers=N_CORES, 
            initializer=initializer_worker, 
            initargs=(y, n_samples_main, n_actuators_main) # Teruskan variabel besar di sini
        ) as executor:
            # Fungsi evaluate_fitness sekarang menggunakan y_global, n_samples_global, n_actuators_global 
            # yang diinisialisasi per proses anak.
            fitness = list(executor.map(evaluate_fitness, population))

        # ... (Seleksi, Crossover, Mutasi, dan Logik GA lainnya tetap sama) ...
        # ... (Logika GA ini tidak perlu proses paralel untuk setiap langkah) ...
        
        new_population = []
        for _ in range(POP_SIZE):
            p1 = selection(population, fitness)
            p2 = selection(population, fitness)
            child = crossover(p1, p2)
            child = mutate(child)
            new_population.append(child)

        population = new_population
        best_idx = np.argmax(fitness)
        print(f"Gen {gen+1:02d} | Best Fitness: {fitness[best_idx]:.4f}")

    # ============================
    # FINAL RESULT
    # ============================
    # Ulangi inisialisasi untuk perhitungan final
    with ProcessPoolExecutor(
        max_workers=N_CORES, 
        initializer=initializer_worker, 
        initargs=(y, n_samples_main, n_actuators_main)
    ) as executor:
        final_fitness = list(executor.map(evaluate_fitness, population))

    best_idx = np.argmax(final_fitness)
    best_rule = population[best_idx]
    best_fit = final_fitness[best_idx]

    # ... (Final printout tetap sama) ...
    print("\n============================")
    print(" BEST GA RULE (FINAL OUTPUT)")
    print("============================")

    for act, val in zip(action_cols, best_rule):
        print(f"{act:35s} = {val}")

    print("\nFinal Fitness:", best_fit)