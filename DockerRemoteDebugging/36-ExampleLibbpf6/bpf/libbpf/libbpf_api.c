#include "libbpf_api.h"
#include "libbpf.h"

int Cbar()
{
    int ncpus = libbpf_num_possible_cpus();
    return ncpus;
}